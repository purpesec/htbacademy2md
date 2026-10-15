package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Args struct {
	moduleUrl   string
	cookies     string
	lang        string
	localImages bool
}

// getArguments parses the command line flags, falling back to values from a
// .env file (./.env or $HOME/.config/htbacademy2md/.env). Explicitly provided
// flags always take precedence over .env values.
func getArguments() Args {
	var mFlag = flag.String("m", "", "(REQUIRED) Academy Module URL to the first page.")
	var cFlag = flag.String("c", "", "Academy Cookies for authorization. Falls back to HTB_COOKIES in .env.")
	var langFlag = flag.String("lang", "", "Language code for translated content (e.g. es, fr, de, ja). Empty = original English content.")
	var imgFlag = flag.Bool("local_images", false, "Save images locally rather than referencing the URL location.")
	var envFlag = flag.String("env", "", "Path to an .env file with default values (default: ./.env or ~/.config/htbacademy2md/.env).")
	flag.Parse()

	// Remember which flags were explicitly set on the command line.
	setFlags := map[string]bool{}
	flag.Visit(func(f *flag.Flag) { setFlags[f.Name] = true })

	env := loadEnv(*envFlag)

	arg := Args{
		moduleUrl:   *mFlag,
		cookies:     *cFlag,
		lang:        *langFlag,
		localImages: *imgFlag,
	}

	if !setFlags["c"] {
		if v := env["HTB_COOKIES"]; v != "" {
			arg.cookies = v
		}
	}
	if !setFlags["lang"] {
		if v := env["HTB_LANG"]; v != "" {
			arg.lang = v
		}
	}
	if !setFlags["local_images"] {
		if v := env["HTB_LOCAL_IMAGES"]; v != "" {
			arg.localImages = parseBool(v)
		}
	}

	if arg.moduleUrl == "" || arg.cookies == "" {
		fmt.Println("Missing required arguments for module URL and HTB Academy Cookies. Please use the -h option for help.")
		fmt.Println("Tip: you can store your cookies and defaults in a .env file (see .env.example).")
		os.Exit(1)
	}

	return arg
}

// loadEnv reads a simple KEY=VALUE file. When path is empty it looks for
// ./.env first and then $HOME/.config/htbacademy2md/.env. Missing files are
// not an error; an explicitly provided path that cannot be read is.
func loadEnv(path string) map[string]string {
	if path != "" {
		env, err := parseEnvFile(path)
		if err != nil {
			die(fmt.Errorf("could not read env file %s: %w", path, err))
		}
		return env
	}

	candidates := []string{".env"}
	if home, err := os.UserHomeDir(); err == nil {
		candidates = append(candidates, filepath.Join(home, ".config", "htbacademy2md", ".env"))
	}

	for _, candidate := range candidates {
		if env, err := parseEnvFile(candidate); err == nil {
			return env
		}
	}

	return map[string]string{}
}

// parseEnvFile reads KEY=VALUE pairs, ignoring blank lines and lines starting
// with '#'. Matching surrounding single or double quotes are stripped.
func parseEnvFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	env := map[string]string{}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024) // cookie values can be long
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			fmt.Printf("Warning: ignoring malformed line %d in %s\n", lineNo, path)
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if len(value) >= 2 {
			first, last := value[0], value[len(value)-1]
			if (first == '"' && last == '"') || (first == '\'' && last == '\'') {
				value = value[1 : len(value)-1]
			}
		}

		env[key] = value
	}

	return env, scanner.Err()
}

// parseBool accepts the usual truthy strings used in config files.
func parseBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "1", "true", "yes", "y", "on":
		return true
	}
	return false
}
