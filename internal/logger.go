package internal

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Debug(msg string)
	DebugWithFields(msg string, fields map[string]interface{})
	Info(msg string)
	InfoWithFields(msg string, fields map[string]interface{})
	Warn(msg string)
	WarnWithFields(msg string, fields map[string]interface{})
	Error(msg string, err error)
	ErrorWithFields(msg string, err error, fields map[string]interface{})
	Fatal(msg string, err error)
	FatalWithFields(msg string, err error, fields map[string]interface{})
}

// Logger representa uma instância do logger
type Logger struct {
	log zerolog.Logger
}

// Config contém as configurações para o logger
type Config struct {
	// OutputPaths define para onde os logs serão enviados
	// Por exemplo: stdout, arquivo, etc.
	OutputPaths []string

	// Level define o nível mínimo de logs que serão registrados
	Level string

	// PrettyPrint formata os logs para melhor leitura humana
	PrettyPrint bool

	// WithCaller adiciona informação sobre o arquivo e linha onde o log foi gerado
	WithCaller bool
}

// NewLogger cria uma nova instância do logger
func NewLogger(config Config) (*Logger, error) {
	// Configuração do zerolog para usar time.RFC3339 para timestamps
	zerolog.TimeFieldFormat = time.RFC3339

	// Definir o nível mínimo de log
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		// Default para info se o nível fornecido for inválido
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	var writers []io.Writer

	// Configurar saídas
	for _, path := range config.OutputPaths {
		switch path {
		case "stdout":
			if config.PrettyPrint {
				writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
			} else {
				writers = append(writers, os.Stdout)
			}
		case "stderr":
			if config.PrettyPrint {
				writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
			} else {
				writers = append(writers, os.Stderr)
			}
		default:
			// Assume que é um caminho de arquivo
			file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				return nil, err
			}
			writers = append(writers, file)
		}
	}

	// Se não houver saídas configuradas, usar stdout
	if len(writers) == 0 {
		if config.PrettyPrint {
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
		} else {
			writers = append(writers, os.Stdout)
		}
	}

	// Criar logger com múltiplas saídas, se necessário
	var log zerolog.Logger
	if len(writers) == 1 {
		log = zerolog.New(writers[0])
	} else {
		log = zerolog.New(zerolog.MultiLevelWriter(writers...))
	}

	// Adicionar timestamp a todos os logs
	log = log.With().Timestamp().Logger()

	// Adicionar informações do caller se configurado
	if config.WithCaller {
		log = log.With().Caller().Logger()
	}

	return &Logger{log: log}, nil
}

// Debug logs a message at debug level
func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

// DebugWithFields logs a message at debug level with additional fields
func (l *Logger) DebugWithFields(msg string, fields map[string]interface{}) {
	event := l.log.Debug()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Info logs a message at info level
func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

// InfoWithFields logs a message at info level with additional fields
func (l *Logger) InfoWithFields(msg string, fields map[string]interface{}) {
	event := l.log.Info()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Warn logs a message at warn level
func (l *Logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

// WarnWithFields logs a message at warn level with additional fields
func (l *Logger) WarnWithFields(msg string, fields map[string]interface{}) {
	event := l.log.Warn()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Error logs a message at error level
func (l *Logger) Error(msg string, err error) {
	l.log.Error().Err(err).Msg(msg)
}

// ErrorWithFields logs a message at error level with additional fields
func (l *Logger) ErrorWithFields(msg string, err error, fields map[string]interface{}) {
	event := l.log.Error().Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Fatal logs a message at fatal level and then exits with status code 1
func (l *Logger) Fatal(msg string, err error) {
	l.log.Fatal().Err(err).Msg(msg)
}

// FatalWithFields logs a message at fatal level with additional fields and then exits
func (l *Logger) FatalWithFields(msg string, err error, fields map[string]interface{}) {
	event := l.log.Fatal().Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}
