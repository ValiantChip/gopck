package pckmngr

import (
	"os"
)

type PackManager struct {
	name    string
	path    string
	version string

	namspaces map[string]*NameSpace
}

func (p *PackManager) Name() string {
	return p.name
}

func (p *PackManager) Version() string {
	return p.version
}

func NewPackManager(name string, path string, version string) *PackManager {
	return &PackManager{name: name, path: path, version: version, namspaces: make(map[string]*NameSpace)}
}

func (p *PackManager) SetNamespace(name string) *NameSpace {
	p.namspaces[name] = &(NameSpace{name: name, path: p.path + p.Name() + "/", folders: make(map[string]*Folder)})
	return p.namspaces[name]
}

func (p *PackManager) CreatePack() error {
	for _, n := range p.namspaces {
		err := n.createNamespace()
		if err != nil {
			return err
		}
	}
	return nil
}

type NameSpace struct {
	name    string
	path    string
	folders map[string]*Folder
}

func (n *NameSpace) createNamespace() error {
	for _, f := range n.folders {
		err := f.createFolder()
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *NameSpace) SetFolder(name string) *Folder {
	n.folders[name] = &(Folder{name: name, path: n.path, files: make(map[string]*File)})
	return n.folders[name]
}

func (n *NameSpace) Name() string {
	return n.name
}

type Folder struct {
	name string
	path string

	files map[string]*File
}

func (dir *Folder) Name() string {
	return dir.name
}

func (dir *Folder) Path() string {
	return dir.path
}

func (dir *Folder) createFolder() error {
	for _, f := range dir.files {
		err := os.Mkdir(dir.Name(), os.ModeDir)
		if err != nil {
			return err
		}
		err = createFile(f, *dir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dir *Folder) SetFile(name string, content []Nest) *File {
	dir.files[name] = &File{name: name, content: content}
	return dir.files[name]
}

type Nest interface {
	Content() string
}

type File struct {
	name    string
	content []Nest
	path    string
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Content() (result string) {
	for _, n := range f.content {
		result += n.Content()
		result += "\n"
	}
	return
}

func createFile(f *File, dir Folder) error {
	err := os.WriteFile(dir.Path()+dir.Name()+"/"+f.Name(), []byte(f.Content()), 0200)
	if err != nil {
		return err
	}

	return nil
}
