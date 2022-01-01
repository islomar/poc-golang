// +build inmemory

package cmd

import (
  "pragprog.com/rggo/interactiveTools/pomo/pomodoro"
  "pragprog.com/rggo/interactiveTools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
  return repository.NewInMemoryRepo(), nil
}
