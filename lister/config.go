package lister

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joshmedeski/sesh/v2/model"
)

func ConfigKey(name string) string {
	return fmt.Sprintf("config:%s", name)
}

func listConfig(l *RealLister) (model.SeshSessions, error) {
	orderedIndex := make([]string, 0)
	directory := make(model.SeshSessionMap)
	for _, session := range l.config.SessionConfigs {
		if session.Name != "" {
			key := ConfigKey(session.Name)
			orderedIndex = append(orderedIndex, key)
			path, err := l.home.ExpandPath(session.Path)
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
				WindowNames:           session.Windows,
				Alias:                 session.Alias,
				AliasAutoConnect:      session.AliasAutoConnect,
				Tmuxp:                 session.Tmuxp,
			}

			if session.WithChildren {
				entries, err := os.ReadDir(path)
				if err != nil {
					continue
				}
				for _, entry := range entries {
					childPath := filepath.Join(path, entry.Name())
					info, err := os.Stat(childPath)
					if err != nil || !info.IsDir() {
						continue
					}
					child := directory[key]
					child.Name = session.Name + "/" + entry.Name()
					child.Path = childPath
					child.Alias = ""
					child.AliasAutoConnect = false
					childKey := ConfigKey(child.Name)
					orderedIndex = append(orderedIndex, childKey)
					directory[childKey] = child
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
	key := ConfigKey(name)
	sessions, _ := listConfig(l)
	if session, exists := sessions.Directory[key]; exists {
		return session, exists
	} else {
		return model.SeshSession{}, false
	}
}
