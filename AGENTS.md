# Agent Instructions for Hestia

## Project Overview
**Hestia** is a family financial management system named after the Greek goddess of hearth and home. The goal is to create a simple and practical tool for managing household finances without over-engineering.

## Core Philosophy
- **Keep it simple**: Don't over-engineer or be overly ambitious
- **Family-focused**: Designed specifically for household financial management
- **Practical approach**: Focus on real-world utility over complex features

## Development Workflow

### Test-Driven Development (TDD)
**CRITICAL**: Always start by creating a test FIRST when implementing new features. This is a strict requirement for this project.

### Language & Runtime
- **Language**: Go 1.25.0
- **Module**: `hestia`
- **Entry Point**: `main.go`

### Commands & Aliases
The project uses **kudasai** for command aliases:

#### Available Commands
```bash
kudasai start    # Equivalent to: go run main.go
kudasai test     # Equivalent to: go test
```

#### Standard Go Commands
```bash
go run main.go           # Run the application
go build -o hestia       # Build binary
go test ./...           # Run all tests
ginkgo                  # Run tests with Ginkgo (preferred)
```

### Testing Framework
- **Primary**: Ginkgo v2 with Gomega matchers
- **Test Structure**: 
  - `hestia_suite_test.go`: Main test suite setup
  - `basic_test.go`: Basic functionality tests
  - Individual test files should follow `*_test.go` naming

### Dependencies
- **Testing**: `github.com/onsi/ginkgo/v2` and `github.com/onsi/gomega`
- **No external runtime dependencies yet** (keep it minimal)

## Planned Features (Financial Domain)

### Core Financial Features
1. **Income tracking** - Salary and other income sources
2. **Bill management** - Recurring and one-time bills
3. **Bank account monitoring** - Balance tracking across accounts
4. **Investment tracking** - Basic investment portfolio monitoring
5. **Expense categorization** - Organize expenses by category

### Implementation Priority
Features should be implemented in order of practical importance for family financial management.

## Code Standards

### Language Policy
- **All code in English**: Variable names, function names, struct fields, etc.
- **No comments**: Keep code self-documenting without comments
- **Spanish UI only**: User-facing content remains in Spanish via i18n

### File Organization
```
hestia/
├── main.go                    # Application entry point
├── hestia_suite_test.go       # Main test suite
├── handlers/                  # Web handlers package
│   ├── home.go                # Home page handler
│   ├── home_test.go           # Handler tests
│   └── handlers_suite_test.go # Handlers test suite
├── i18n/                      # Internationalization package
│   ├── translations.go        # Translation loading
│   ├── translations_test.go   # Translation tests
│   └── i18n_suite_test.go     # I18n test suite
├── static/                    # Static assets
│   ├── templates/             # HTML templates
│   │   └── index.html
│   └── locales/               # Translation files
│       └── es.json
├── .github/
│   └── workflows/
│       └── test.yml           # CI/CD pipeline
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── .kudasai.json              # Command aliases
├── .gitignore                 # Git ignore rules
├── README.md                  # Project documentation
└── AGENTS.md                  # This file
```

### Communication Style
- **Be concise and direct** - User prefers brief, practical communication
- **Avoid over-explanation** - Focus on actionable information
- **Use this file** instead of creating separate CLAUDE.md files

## Technical Notes

### Current State
- **Status**: Structured web application with modular architecture
- **Tests**: Complete Ginkgo test suite with package-level testing
- **Dependencies**: Minimal - only Ginkgo/Gomega for testing, Pico CSS via CDN
- **CI/CD**: GitHub Actions for automated testing

### Development Reminders
- Always read this AGENTS.md file at the start of any session
- Follow TDD workflow religiously 
- Keep the scope practical and family-oriented
- Prefer Go standard library over external dependencies when possible
- Test coverage should be comprehensive but not obsessive

## Linting & Code Quality
- **Linting**: [To be determined - check for existing linting setup]
- **Code Style**: Follow standard Go conventions and gofmt formatting