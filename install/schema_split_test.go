package install

import (
	"strings"
	"testing"
)

func TestSplitSQLStatements_PreservesSemicolonInComment(t *testing.T) {
	stmts := splitSQLStatements(string(schemaSQL))
	if len(stmts) != 8 {
		t.Fatalf("want 8 statements, got %d", len(stmts))
	}
	for i, stmt := range stmts {
		if !strings.HasPrefix(stmt, "CREATE TABLE") {
			t.Fatalf("stmt %d does not start with CREATE TABLE: %q", i, truncate(stmt, 40))
		}
		if strings.Count(stmt, "'")%2 != 0 {
			t.Fatalf("stmt %d has unbalanced quotes", i)
		}
	}
	dept := stmts[0]
	if !strings.Contains(dept, "COMMENT '状态(1:正常;0:禁用)'") {
		t.Fatalf("sys_dept status COMMENT was split incorrectly:\n%s", dept)
	}
}
