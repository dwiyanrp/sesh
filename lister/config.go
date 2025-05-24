package lister

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joshmedeski/sesh/v2/model"
)

func configKey(name string) string {
	return fmt.Sprintf("config:%s", name)
}

func listConfig(l *RealLister) (model.SeshSessions, error) {
	orderedIndex := make([]string, 0)
	directory := make(model.SeshSessionMap)
	for _, session := range l.config.SessionConfigs {
		if session.Name != "" {
			key := configKey(session.Name)
			orderedIndex = append(orderedIndex, key)
			path, err := l.home.ExpandHome(session.Path)
			if err != nil {
				return model.SeshSessions{}, fmt.Errorf("couldn't expand home: %q", err)
			}

			if session.StartupCommand != "" && session.DisableStartCommand {
				return model.SeshSessions{}, fmt.Errorf("startup_command and disable_start_command are mutually exclusive")
			}

			directory[key] = model.SeshSession{
				Src:                   "config",
				Name:                  session.Name,
				Path:                  path,
				StartupCommand:        session.StartupCommand,
				PreviewCommand:        session.PreviewCommand,
				DisableStartupCommand: session.DisableStartCommand,
				Tmuxinator:            session.Tmuxinator,
			}

			if session.WithChild {
				// Read directory entries
				entries, err := os.ReadDir(path)
				if err != nil {
					return model.SeshSessions{}, fmt.Errorf("couldn't read dir: %q", err)
				}

				// Add session for each directory entry
				for _, entry := range entries {
					// Check if the entry is a directory or a symlink to a directory
					if entry.IsDir() || (entry.Type()&os.ModeSymlink != 0) {
						targetPath := filepath.Join(path, entry.Name())

						// If it's a symlink, resolve the actual path
						if entry.Type()&os.ModeSymlink != 0 {
							resolvedPath, err := filepath.EvalSymlinks(targetPath)
							if err != nil {
								fmt.Printf("Warning: could not resolve symlink %q: %v\n", targetPath, err)
								continue // Skip this entry if symlink resolution fails
							}

							// Check if the resolved path is a directory
							resolvedInfo, err := os.Stat(resolvedPath)
							if err != nil {
								fmt.Printf("Warning: could not stat resolved path %q: %v\n", resolvedPath, err)
								continue // Skip if stat fails
							}
							if !resolvedInfo.IsDir() {
								fmt.Printf("Warning: resolved symlink %q is not a directory\n", resolvedPath)
								continue // Skip if not a directory
							}
							targetPath = resolvedPath
						}

						sName := session.Name + "/" + entry.Name()
						key := configKey(sName)
						orderedIndex = append(orderedIndex, key)
						directory[key] = model.SeshSession{
							Src:                   "config",
							Name:                  sName,
							Path:                  targetPath, // Use resolved path for symlinks
							StartupCommand:        session.StartupCommand,
							PreviewCommand:        session.PreviewCommand,
							DisableStartupCommand: session.DisableStartCommand,
							Tmuxinator:            session.Tmuxinator,
						}
					}
				}
			}
		}
	}
	return model.SeshSessions{
		Directory:    directory,
		OrderedIndex: orderedIndex,
	}, nil
}

func (l *RealLister) FindConfigSession(name string) (model.SeshSession, bool) {
	key := configKey(name)
	sessions, _ := listConfig(l)
	if session, exists := sessions.Directory[key]; exists {
		return session, exists
	} else {
		return model.SeshSession{}, false
	}
}
