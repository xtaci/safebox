package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func showDirWindow(inputField *tview.InputField) {
	const (
		windowName   = "showDirWindow"
		windowWidth  = 80
		windowHeight = 30
		windowTitle  = S_WINDOW_SHOWDIR_TITLE
	)

	rootDir := "/"
	if runtime.GOOS == "windows" {
		rootDir = "C:\\"
	}

	node := tview.NewTreeNode(rootDir)
	tree := tview.NewTreeView().
		SetRoot(node).
		SetCurrentNode(node)

	tree.SetBorder(true).SetTitle(windowTitle)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorDarkGreen)
			}
			target.AddChild(node)
		}
	}

	add(node, rootDir)

	// path route
	var route []string
	// recursive to root
	dir := inputField.GetText()

	// get file info
	fi, _ := os.Stat(dir)

	for dir != rootDir {
		if dir == "." {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			dir = wd
		}

		dir = filepath.Dir(dir)
		route = append(route, dir)
	}

	// expand the tree to path the inputfield
	prevNode := node
	for i := len(route) - 1; i >= 0; i-- {
		children := prevNode.GetChildren()
		for _, children := range children {
			dir := children.GetReference().(string)
			if dir == route[i] {
				tree.SetCurrentNode(children)
				files, err := ioutil.ReadDir(dir)
				if err != nil {
					panic(err)
				}

				for _, file := range files {
					node := tview.NewTreeNode(file.Name()).
						SetReference(filepath.Join(route[i], file.Name())).
						SetSelectable(true)
					if file.IsDir() {
						node.SetColor(tcell.ColorDarkGreen)
					} else {
						if os.SameFile(fi, file) { // select the given file if match
							tree.SetCurrentNode(node)
						}
					}
					children.AddChild(node)
				}

				children.Expand()
				prevNode = children
				break
			}
		}
	}

	// Change input field text based on tree selection
	tree.SetChangedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		path := reference.(string)
		// set this path to input text
		inputField.SetText(path)
	})

	// If a directory was selected, open it.
	// If a file was selected, close the page
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		path := reference.(string)

		fi, err := os.Stat(inputField.GetText())
		if err != nil {
			panic(err)
		}

		if fi.IsDir() { // check if the selected path is dir
			if len(children) == 0 {
				// Load and show files in this directory.
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}
		} else { // file selected, close this page
			layoutRoot.RemovePage(windowName)
		}
	})

	// set input field to current node
	currentNode := tree.GetCurrentNode()
	reference := currentNode.GetReference()
	if reference != nil {
		inputField.SetText(reference.(string))
	}

	layoutRoot.AddPage(windowName, popup(windowWidth, windowHeight, tree), true, true)
}
