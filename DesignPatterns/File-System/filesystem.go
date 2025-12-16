package filesystem

import (
	"fmt"
)

type Node interface {
	GetName() string
	GetSize() int
	Print(indent string)
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Print(indent string) {
	fmt.Printf("%s📄 %s (%dKB)\n", indent, f.name, f.size)
}

type Directory struct {
	name     string
	children []Node
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) Add(node Node) {
	d.children = append(d.children, node)
}

func (d *Directory) Remove(node Node) {
	for i, child := range d.children {
		if child == node {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) GetSize() int {
	total := 0
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}

func (d *Directory) Print(indent string) {
	fmt.Printf("%s📁 %s (%dKB)\n", indent, d.name, d.GetSize())
	for _, child := range d.children {
		child.Print(indent + "  ")
	}
}

func main() {
	fmt.Println("File System Composite Pattern Example")
	// root := filesystem.NewDirectory("root")

	// file1 := filesystem.NewFile("a.txt", 10)
	// file2 := filesystem.NewFile("b.txt", 20)

	// subDir := filesystem.NewDirectory("docs")
	// subDir.Add(filesystem.NewFile("readme.md", 5))

	// root.Add(file1)
	// root.Add(file2)
	// root.Add(subDir)

	// root.Print("")
}
