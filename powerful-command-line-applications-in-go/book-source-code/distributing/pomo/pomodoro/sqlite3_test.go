//+build !inmemory

package pomodoro_test

import (
  "io/ioutil"
  "os"
  "testing"

  "pragprog.com/rggo/interactiveTools/pomo/pomodoro"
  "pragprog.com/rggo/interactiveTools/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
  t.Helper()

  tf, err := ioutil.TempFile("", "pomo")
  if err != nil {
    t.Fatal(err)
  }
  tf.Close()

  dbRepo, err := repository.NewSQLite3Repo(tf.Name())

  if err != nil {
    t.Fatal(err)
  }

  return dbRepo, func() {
    os.Remove(tf.Name())
  }
}
