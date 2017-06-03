package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func possession(name string) string {
	if strings.HasSuffix(name, "s") || strings.HasSuffix(name, "z") {
		return fmt.Sprintf("%s'", name)
	}

	return fmt.Sprintf("%s's", name)
}

func diariesDirectory() (string, error) {
	dir, err := exeDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "diaries"), nil
}

func exePath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("%s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("%s is directory", p)
		}
	}
	return "", err
}

func exeDirectory() (string, error) {
	p, err := exePath()
	if err != nil {
		return "", err
	}

	return filepath.Dir(p), nil
}

func matchProfile(profile string, profiles []string) (string, error) {
	tmp := strings.ToLower(profile)

	// try lowercased exact-matches
	for _, candidate := range profiles {
		lower := strings.ToLower(candidate)

		if profile == lower {
			return candidate, nil
		}
	}

	// try globbing
	candidates := make([]string, 0)
	pattern := fmt.Sprintf("%s*", tmp)

	for _, candidate := range profiles {
		lower := strings.ToLower(candidate)

		match, err := filepath.Match(pattern, lower)
		if err == nil && match {
			candidates = append(candidates, candidate)
		}
	}

	if len(candidates) == 1 {
		return candidates[0], nil
	} else if len(candidates) > 1 {
		return "", fmt.Errorf("The savegame name '%s' is ambiguous, could mean any of [%s].", profile, strings.Join(profiles, ", "))
	}

	return "", fmt.Errorf("Could not find a savegame called '%s'.", profile)
}
