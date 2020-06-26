package main

import (
	"testing"
)

func TestUppercase(t *testing.T) {
	owner, repo := getOwnerAndRepo("LeCompany/importantcode")
	if owner != "LeCompany" {
		t.Errorf("owner = %s; want LeCompany", owner)
	}
	if repo != "importantcode" {
		t.Errorf("repo = %s; want importantcode", repo)
	}
}

func TestHyphen(t *testing.T) {
	owner, repo := getOwnerAndRepo("company/secret-code")
	if owner != "company" {
		t.Errorf("owner = %s; want company", owner)
	}
	if repo != "secret-code" {
		t.Errorf("repo = %s; want secret-code", repo)
	}
}
