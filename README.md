# 🏛️ Hestia

**Family financial management system** - Simple and practical household finance tracking.

Named after the Greek goddess of hearth and home, Hestia helps families manage their finances without over-engineering.

## ✨ Features

- 📊 **Income tracking** - Monitor salary and other income sources
- 💰 **Expense management** - Track and categorize household expenses
- 🏦 **Bank account monitoring** - Keep tabs on account balances
- 📈 **Investment tracking** - Basic investment portfolio monitoring
- 🌐 **Spanish interface** - Family-friendly Spanish UI
- 🌙 **Dark mode** - Automatic light/dark theme switching

## 🚀 Getting Started

### Prerequisites

- Go 1.25.0 or higher
- [kudasai](https://github.com/vibaiher/kudasai) (optional, for command aliases)

### Installation

```bash
git clone https://github.com/vibaiher/hestia.git
cd hestia
go mod download
```

### Running the Application

```bash
# Using kudasai (recommended)
kudasai start

# Or using Go directly
go run main.go
```

Visit [http://localhost:8080](http://localhost:8080) to access Hestia.

### Running Tests

```bash
# Using kudasai
kudasai test

# Or using Go directly
go test ./...

# With Ginkgo for better output
ginkgo
```

## 🏗️ Architecture

Hestia follows Go best practices with a modular package structure:

```
hestia/
├── main.go              # Application entry point
├── handlers/            # Web request handlers
├── i18n/               # Internationalization
├── static/             # Templates and assets
└── *_test.go          # Comprehensive test suite
```

## 🧪 Testing

- **Framework**: Ginkgo v2 with Gomega matchers
- **Philosophy**: Test-driven development (TDD)
- **Coverage**: Unit tests for all packages
- **CI/CD**: Automated testing via GitHub Actions

## 🎨 Design

- **CSS Framework**: Pico CSS (lightweight, semantic)
- **Philosophy**: Keep it simple, family-focused
- **Responsive**: Works on desktop and mobile
- **Accessibility**: Semantic HTML with proper ARIA

## 🤝 Contributing

This project follows the principle of simplicity over complexity. When contributing:

1. Write tests first (TDD)
2. Keep code in English, UI in Spanish
3. Follow Go conventions
4. Maintain minimal dependencies

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🏠 Philosophy

> "The hearth is the heart of the home" - Focus on practical utility over complex features.