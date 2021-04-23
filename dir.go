package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func showDirWindow(inputField *tview.InputField) {
	windowName := "showDirWindow"
	rootDir := "/"

	node := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(node).
		SetCurrentNode(node)

	tree.SetBorder(true).SetTitle("Directory")

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
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	add(node, rootDir)

	// path route
	var route []string
	// recursive to root
	dir := inputField.GetText()

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

	// create a path to inputfield
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
						SetSelectable(file.IsDir())
					if file.IsDir() {
						node.SetColor(tcell.ColorGreen)
					}
					children.AddChild(node)
				}

				children.Expand()
				prevNode = children
				break
			}
		}
	}

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		path := reference.(string)
		inputField.SetText(path + "/.safebox.key")

		if len(children) == 0 {
			// Load and show files in this directory.
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	root.AddPage(windowName, popup(40, 40, tree), true, true)
}
