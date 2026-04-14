# 🚀 Glask - AI-Powered Terminal Integration

**Glask** é um assistente de IA integrado ao Terminal Linux que evoluirá de um simples chat para um agente autônomo completo, capaz de executar comandos, gerenciar arquivos e otimizar workflows de desenvolvimento.

## 📋 Sobre o Projeto

### Visão
Transformar a forma como desenvolvedores interagem com o terminal, fornecendo um assistente de IA conversacional que evolui para um agente autônomo integrado, compreendendo contexto, executando tarefas complexas e aprendendo com o histórico de interações.

### Objetivo Geral
Criar uma ferramenta CLI que:
- 💬 Funcione como um chat interativo com IA
- 🤖 Evolua para um agente autônomo
- 📁 Integre-se perfeitamente com o sistema de arquivos
- 🔌 Suporte plugins e extensões
- 📊 Mantenha histórico e contexto das interações

---

## 🎯 Fases de Desenvolvimento

### **Fase 1: Chat Interativo (Atual)**
- [ ] Interface CLI básica com Cobra
- [ ] Integração com API de IA (OpenAI/Claude/Ollama)
- [ ] Chat simples via terminal
- [ ] Histórico de conversas em arquivo
- [ ] Logs estruturados de interações
- [ ] Temas e personalização básica

### **Fase 2: Contexto e Awareness**
- [ ] Leitura de contexto do projeto (arquivos, estrutura)
- [ ] Entendimento de pwd e estrutura de diretórios
- [ ] Análise de código nos diretórios
- [ ] Memória de contexto entre sessões
- [ ] Configuração por projeto

### **Fase 3: Execução de Comandos**
- [ ] Sugestão segura de comandos
- [ ] Execução de comandos com aprovação
- [ ] Sandbox para testes
- [ ] Histórico de comandos executados
- [ ] Rollback de operações

### **Fase 4: Agente Autônomo**
- [ ] Planejamento de tarefas
- [ ] Execução automática de workflows
- [ ] Monitoramento e diagnóstico
- [ ] Resolução automática de problemas comuns
- [ ] Otimização de código

### **Fase 5: Extensibilidade**
- [ ] Sistema de plugins
- [ ] Hooks de eventos
- [ ] Integração com ferramentas externas
- [ ] API REST para integração

---

## ✨ Características Planejadas

### Chat Interativo
```bash
$ glask chat
> Qual é a melhor forma de debugar um memory leak em Go?
> Crie um arquivo JSON com a estrutura de um usuário
> Me mostre os 10 arquivos mais recentemente modificados
```

### Execução de Comandos Sugerida
```bash
$ glask exec "comprimir todos os arquivos .log"
[IA Suggestion] Você deseja executar:
tar -czf logs-backup.tar.gz *.log
[Aprovar? y/n]
```

### Agente de Background
```bash
$ glask agent start --watch .
Monitorando alterações nos arquivos...
[Agent] Detectei um novo arquivo de teste
[Agent] Executando testes...
```

---

## 📦 Stack Tecnológico

### Dependências Principais

#### **Flags e CLI**
- **Cobra** - Framework CLI profissional
- **Viper** - Gerenciamento de configuração

#### **UI/UX**
- **Bubble Tea** - Framework para TUIs interativas
- **Lipgloss** - Estilização de componentes
- **Bubbles** - Componentes reutilizáveis

#### **Logs Estruturados**
- **Zap** - Logger de alto performance
- **go.uber.org/zap** - Logs estruturados

#### **IA Integration**
- **openai-go** ou similar - Integração com APIs de IA

### Instalação de Dependências

```bash
# Dependências principais
go get -u github.com/spf13/cobra@latest
go get -u github.com/spf13/viper
go get -u go.uber.org/zap
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
```

---

## ## Bibliotecas Recomendadas

### 📊 Flags e Comandos CLI

#### **Cobra** (Recomendado)
- **Descrição**: Framework completo para criar CLIs profissionais com comandos, subcomandos e flags
- **Instalação**: `go get -u github.com/spf13/cobra@latest`
- **Uso**:
  ```go
  import "github.com/spf13/cobra"
  
  var rootCmd = &cobra.Command{
    Use:   "glask",
    Short: "Glask CLI application",
    Run: func(cmd *cobra.Command, args []string) {
      // código do comando
    },
  }
  ```
- **Vantagens**: 
  - Suporte completo a subcomandos
  - Geração automática de help e completion
  - Integração com Viper para configuração

#### **Viper** (Complementar)
- **Descrição**: Gerenciador de configuração que trabalha perfeitamente com Cobra
- **Instalação**: `go get github.com/spf13/viper`
- **Uso**: Carrega configs de arquivos, env vars e flags
- **Vantagens**:
  - Suporte a múltiplos formatos (YAML, JSON, TOML, HCL)
  - Integração perfeita com Cobra
  - Watch de mudanças em arquivos de config

#### **pflag** (Alternativa leve)
- **Descrição**: Implementação POSIX-style de flags
- **Instalação**: `go get github.com/spf13/pflag`
- **Uso**: Base do Cobra, pode ser usado isoladamente

---

### 🎨 Interface CLI / UI

#### **Bubble Tea** (Recomendado para TUIs interativas)
- **Descrição**: Framework para criar interfaces de terminal sofisticadas e interativas
- **Instalação**: `go get github.com/charmbracelet/bubbletea`
- **Uso**:
  ```go
  import tea "github.com/charmbracelet/bubbletea"
  
  type Model struct {
    // estado da aplicação
  }
  
  func (m Model) Init() tea.Cmd { return nil }
  func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
  func (m Model) View() string { return "Conteúdo" }
  ```
- **Vantagens**:
  - Padrão elegante baseado em Elm
  - Suporte a eventos de mouse e teclado
  - Rendering eficiente

#### **Lipgloss** (Estilização com Bubble Tea)
- **Descrição**: Biblioteca para criar UIs estilizadas com cores, layouts e bordas
- **Instalação**: `go get github.com/charmbracelet/lipgloss`
- **Uso**: Styling de componentes da UI
- **Vantagens**:
  - Fácil criar layouts responsivos
  - Tema de cores customizável
  - Renderização otimizada

#### **Glamour** (Renderização de Markdown)
- **Descrição**: Renderiza Markdown em terminal com estilo
- **Instalação**: `go get github.com/charmbracelet/glamour`
- **Uso**: Exibir conteúdo Markdown formatado no terminal

#### **Progressbar** (Alternativa simples)
- **Descrição**: Barra de progresso simples para CLI
- **Instalação**: `go get github.com/schollz/progressbar/v3`
- **Uso**: Feedback visual de operações longas

#### **Spinner** (Indicador de progresso)
- **Descrição**: Spinners animados para operações em progresso
- **Instalação**: `go get github.com/briandowns/spinner`
- **Uso**: Indicador visual durante processamento

---

### 📝 Logs Estruturados

#### **Zap** (Recomendado - Alto Performance)
- **Descrição**: Logger estruturado extremamente rápido e flexível
- **Instalação**: `go get -u go.uber.org/zap`
- **Uso**:
  ```go
  import "go.uber.org/zap"
  
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  
  logger.Info("Mensagem de informação",
    zap.String("user", "João"),
    zap.Int("count", 42),
  )
  ```
- **Vantagens**:
  - Extremamente rápido (structured logging)
  - Suporte a múltiplos níveis de log
  - Fácil configuração e formatação
  - Integração com honeycomb, datadog, etc

#### **Logrus** (Alternativa popular)
- **Descrição**: Logger estruturado com interface simples
- **Instalação**: `go get github.com/sirupsen/logrus`
- **Uso**:
  ```go
  import log "github.com/sirupsen/logrus"
  
  log.WithFields(log.Fields{
    "event": "login",
    "topic": "authentication",
  }).Info("Login realizado")
  ```
- **Vantagens**:
  - API intuitiva e familiar
  - Hooks para processamento customizado
  - Formatação em JSON disponível

#### **Slog** (Padrão da stdlib - Go 1.21+)
- **Descrição**: Logger estruturado nativo do Go 1.21+
- **Uso**:
  ```go
  import "log/slog"
  
  slog.Info("Mensagem", "user", "João", "count", 42)
  ```
- **Vantagens**:
  - Nenhuma dependência externa
  - Performance excelente
  - API limpa e consistente

#### **Log4go** (Logging avançado)
- **Descrição**: Logger inspirado em Log4j do Java
- **Instalação**: `go get github.com/ksambright/log4go`
- **Uso**: Configuração complexa de logs com múltiplos appenders

---

## Stack Recomendado Completo

### Para um CLI simples com logs:
```bash
go get -u github.com/spf13/cobra@latest
go get -u go.uber.org/zap
```

### Para um CLI completo:
```bash
go get -u github.com/spf13/cobra@latest
go get -u github.com/spf13/viper
go get -u go.uber.org/zap
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
```

### Para um TUI (Terminal User Interface) interativo:
```bash
go get -u github.com/spf13/cobra@latest
go get -u github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
go get -u go.uber.org/zap
```

---

## Estrutura do Projeto

```
Glask/
├── cmd/
│   └── glask/
│       └── main.go           # Ponto de entrada
├── internal/
│   ├── ai/                   # Lógica de IA
│   ├── config/               # Configurações
│   └── ui/                   # Componentes de UI/CLI
├── pkg/                      # Pacotes reutilizáveis
├── scripts/                  # Scripts utilitários
├── go.mod
└── README.md
```

---

## Exemplos Rápidos

### Exemplo 1: CLI com Cobra + Zap

```go
package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var rootCmd = &cobra.Command{
		Use:   "glask",
		Short: "Glask - CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Comando executado com sucesso")
		},
	}

	rootCmd.Execute()
}
```

### Exemplo 2: Logs Estruturados com Zap

```go
logger.Info("Operação realizada",
	zap.String("operacao", "criar_usuario"),
	zap.String("usuario", "joao@example.com"),
	zap.Int("tempo_ms", 156),
	zap.Bool("sucesso", true),
)
```

### Exemplo 3: TUI com Bubble Tea

```go
package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	choice int
	chosen bool
}

func (m Model) Init() tea.Cmd   { return nil }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
func (m Model) View() string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("Olá Glask!")
}

func main() {
	p := tea.NewProgram(Model{})
	p.Run()
}
```

---

## Comparação de Loggers

| Característica | Zap | Logrus | Slog |
|---|---|---|---|
| Performance | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Facilidade de uso | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| Dependências | 1 | 2 | 0 |
| Estruturado | Sim | Sim | Sim |
| JSON Output | Sim | Sim | Sim |
| Customização | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |

---

## Referências

- [Cobra Documentation](https://cobra.dev/)
- [Zap Documentation](https://pkg.go.dev/go.uber.org/zap)
- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Viper Documentation](https://github.com/spf13/viper)
- [Lipgloss Documentation](https://github.com/charmbracelet/lipgloss)

---

---

## 🏗️ Arquitetura do Projeto

```
Glask/
├── cmd/
│   └── glask/
│       └── main.go              # Ponto de entrada
├── internal/
│   ├── ai/                      # Integração com modelos de IA
│   │   ├── client.go            # Cliente de IA (OpenAI/Claude/etc)
│   │   ├── context.go           # Gerenciamento de contexto
│   │   └── models.go            # Modelos de resposta
│   ├── config/                  # Configuração da aplicação
│   │   ├── config.go            # Carregamento de config
│   │   └── constants.go         # Constantes
│   ├── ui/                      # Interface de usuário
│   │   ├── chat.go              # Chat interativo
│   │   ├── styles.go            # Temas e estilos
│   │   └── components.go        # Componentes reutilizáveis
│   ├── logger/                  # Logging estruturado
│   │   └── logger.go            # Setup do Zap
│   ├── commands/                # Implementação de comandos
│   │   ├── chat.go              # Comando chat
│   │   ├── exec.go              # Comando exec
│   │   ├── agent.go             # Comando agent
│   │   └── history.go           # Comando history
│   ├── history/                 # Gerenciamento de histórico
│   │   └── storage.go           # Persistência de histórico
│   └── agent/                   # Lógica do agente autônomo
│       ├── planner.go           # Planejamento de tarefas
│       ├── executor.go          # Execução de tarefas
│       └── monitor.go           # Monitoramento
├── pkg/                         # Pacotes reutilizáveis
│   ├── terminal/                # Utilitários de terminal
│   └── file/                    # Utilitários de arquivo
├── docs/                        # Documentação
├── scripts/                     # Scripts de desenvolvimento
├── .env.example                 # Exemplo de configuração
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Como Começar

### Pré-requisitos
- Go 1.21 ou superior
- Linux/macOS (Windows com WSL)
- Chave de API de IA (OpenAI, Claude, etc.) - para Fase 2+

### Instalação

1. **Clone o repositório**
```bash
git clone https://github.com/seu-usuario/glask.git
cd glask
```

2. **Instale as dependências**
```bash
go mod download
```

3. **Copie o arquivo de configuração**
```bash
cp .env.example .env
# Edite .env com suas chaves de API
```

4. **Compile a aplicação**
```bash
go build -o glask ./cmd/glask
```

5. **Execute**
```bash
./glask chat
```

---

## 💻 Exemplos de Uso

### Chat Simples
```bash
$ glask chat
Loading Glask...

> Como faço para listar arquivos recursivamente em Go?
Glask: Você pode usar a função filepath.Walk() da biblioteca padrão. 
Aqui está um exemplo...

> Pode criar um snippet para mim?
[Arquivo criado: snippet_filepath.go]
```

### Consultar Histórico
```bash
glask history --last 5
```

### Ver Configurações
```bash
glask config show
```

### Definir Modelo de IA
```bash
glask config set model gpt-4
```

---

## 🔧 Configuração

### Arquivo `.env`
```bash
# Chave de API
OPENAI_API_KEY=sk-...
AI_MODEL=gpt-4
AI_PROVIDER=openai  # openai, claude, ollama, etc

# Logging
LOG_LEVEL=info  # debug, info, warn, error
LOG_FILE=./logs/glask.log

# UI
UI_THEME=dark  # dark, light, auto
ENABLE_COLORS=true

# Histórico
HISTORY_DIR=~/.glask/history
MAX_HISTORY=1000

# Agent
AGENT_ENABLED=false
AGENT_AUTO_EXECUTE=false
AGENT_SANDBOX=true
```

### Configuração por Projeto
Crie um arquivo `.glaskrc.yaml` na raiz do projeto:
```yaml
name: "Meu Projeto"
description: "Descrição do projeto"
ai:
  model: gpt-4
  temperature: 0.7
  system_prompt: "Você é um assistente de desenvolvimento especializado em Go"
history:
  enabled: true
  file: ./glask-history.json
agent:
  enabled: false
  auto_execute: false
```

---

## 📊 Diagrama de Fluxo

### Chat Simples
```
User Input
    ↓
Parse Command
    ↓
Load Context (files, pwd, etc)
    ↓
Send to AI API
    ↓
Format Response
    ↓
Display & Log
    ↓
Save to History
```

### Execução de Comando
```
User Request
    ↓
AI Suggests Command
    ↓
Show Suggestion
    ↓
User Approves?
    ├→ Yes: Execute in Sandbox
    └→ No: Discard
    ↓
Log Result
```

### Agent Autônomo (Fase 4)
```
Trigger Event
    ↓
Agent Analyzes Context
    ↓
Plans Tasks
    ↓
Executes (if safe)
    ↓
Monitors Results
    ↓
Learns & Adapts
```

---

## 🔐 Segurança

- ✅ Execução de comandos requer aprovação manual (inicialmente)
- ✅ Sandbox para testes antes de execução real
- ✅ Whitelist de comandos permitidos (Fase 3+)
- ✅ Audit log completo de todas as ações
- ✅ Integração com sistema de permissões Linux
- ✅ Detecção de comandos perigosos

---

## 📝 Logging

Todos os eventos são registrados estruturadamente com Zap:

```json
{
  "timestamp": "2024-04-14T10:30:45Z",
  "level": "info",
  "logger": "glask",
  "message": "Chat message processed",
  "user_input_length": 42,
  "ai_model": "gpt-4",
  "response_time_ms": 1250,
  "tokens_used": 315
}
```

Logs disponíveis em:
- Console (terminal)
- Arquivo: `~/.glask/logs/glask.log`

---

## 🧪 Testes

```bash
# Rodar todos os testes
go test ./...

# Com coverage
go test -cover ./...

# Testes específicos
go test -run TestChatCommand ./...
```

---

## 📚 Documentação Adicional

- [Arquitetura Detalhada](./docs/ARCHITECTURE.md)
- [Guia de Contribuição](./docs/CONTRIBUTING.md)
- [API de IA](./docs/AI_API.md)
- [Roadmap Completo](./docs/ROADMAP.md)

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📄 Licença

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

---

## 👨‍💻 Autor

**Heitor** - [@heitor-hsantos](https://github.com/heitor-hsantos)

---

## 🌟 Roadmap

### Q2 2026 (Próximo)
- [x] Setup inicial do projeto
- [ ] Implementar chat básico
- [ ] Integração com OpenAI
- [ ] Histórico de conversas

### Q3 2026
- [ ] Contexto de projeto
- [ ] Leitura de arquivos
- [ ] Sugestão de comandos

### Q4 2026
- [ ] Execução de comandos com aprovação
- [ ] Sandbox para testes
- [ ] Sistema de plugins

### 2027
- [ ] Agente autônomo
- [ ] Machine learning local
- [ ] Integração com Git/GitHub
- [ ] Integração com CI/CD

---

## 💡 Ideias Futuras

- [ ] Integração com Docker
- [ ] Suporte a múltiplos idiomas
- [ ] Cliente web para Glask
- [ ] Sync entre múltiplas máquinas
- [ ] Integração com Slack/Discord
- [ ] Video tutorials & demos

---

## 🆘 Support

Encontrou um bug? Tem uma sugestão? Abra uma [issue](https://github.com/seu-usuario/glask/issues)!

---

## 🎉 Agradecimentos

- [Charmbracelet](https://github.com/charmbracelet) - UI components
- [Spf13](https://github.com/spf13) - Cobra & Viper
- [Uber](https://github.com/uber-go) - Zap logger

---

**Última atualização**: Abril de 2024
**Status**: 🔨 Em desenvolvimento - Fase 1

