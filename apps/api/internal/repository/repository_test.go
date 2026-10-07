package repository_test

import (
	"testing"

	"github.com/codeatlas/api/internal/repository"
	"github.com/codeatlas/api/internal/testutil"
)

func setup(t *testing.T) *repository.Service {
	t.Helper()
	db, err := testutil.NewTestDB()
	if err != nil {
		t.Skipf("skipping: no test db: %v", err)
	}
	return repository.NewService(db)
}

func TestCreate(t *testing.T) {
	svc := setup(t)
	repo, err := svc.Create(1, "my-repo", "/data/repos/1/my-repo")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if repo.ID == 0 {
		t.Fatal("expected non-zero ID")
	}
	if repo.ProjectID != 1 || repo.Name != "my-repo" {
		t.Fatalf("unexpected repo: %+v", repo)
	}
}

func TestList(t *testing.T) {
	svc := setup(t)
	if _, err := svc.Create(2, "repo-a", "/data/repos/2/repo-a"); err != nil {
		t.Fatalf("Create repo-a: %v", err)
	}
	if _, err := svc.Create(2, "repo-b", "/data/repos/2/repo-b"); err != nil {
		t.Fatalf("Create repo-b: %v", err)
	}

	repos, err := svc.List(2)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(repos) < 2 {
		t.Fatalf("expected at least 2 repos, got %d", len(repos))
	}
}

func TestUpload(t *testing.T) {
	t.Skip("implement with testutil.NewTestDB()")
}
