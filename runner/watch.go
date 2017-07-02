package runner

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dnephin/filewatcher/files"
	"gopkg.in/fsnotify.v1"
)

// Watch for events from the watcher and handle them with the runner
func Watch(watcher *fsnotify.Watcher, runner *Runner) error {
	for {
		select {
		case event := <-watcher.Events:
			log.Debugf("Event: %s", event)

			if isNewDir(event, runner.excludes) {
				log.Debugf("Watching new directory: %s", event.Name)
				watcher.Add(event.Name)
				continue
			}

			runner.HandleEvent(event)

		case err := <-watcher.Errors:
			return err
		}
	}
}

func isNewDir(event fsnotify.Event, exclude *files.ExcludeList) bool {
	if event.Op&fsnotify.Create != fsnotify.Create {
		return false
	}

	fileInfo, err := os.Stat(event.Name)
	if err != nil {
		log.Warnf("Failed to stat %s: %s", event.Name, err)
		return false
	}

	return fileInfo.IsDir() && !exclude.IsMatch(event.Name)
}