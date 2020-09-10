package infrastructure_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/crypto-com/chainindex/infrastructure"
	"github.com/crypto-com/chainindex/usecase"
)

const RFC3339_TIME_REGEX = `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`

var _ = Describe("Zerolog", func() {
	It("should implement Logger interface", func() {
		mockWriter := NewMockWriter()
		var _ usecase.Logger = infrastructure.NewZerologLogger(mockWriter)
	})

	Describe("SetLogLevel", func() {
		It("should log only those with levels higher than set level", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("any info message"))

			logger.SetLogLevel(usecase.LOG_LEVEL_ERROR)

			logger.Info("any info message")
			Expect(mockWriter).NotTo(gbytes.Say("any info message"))

			logger.Error("any error message")
			Expect(mockWriter).To(gbytes.Say("any error message"))
		})

		It("should log nothing when disabled", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.SetLogLevel(usecase.LOG_DISABLED)

			logger.Debug("any debug message")
			Expect(mockWriter).NotTo(gbytes.Say("any debug message"))

			logger.Info("any info message")
			Expect(mockWriter).NotTo(gbytes.Say("any info message"))

			logger.Error("any error message")
			Expect(mockWriter).NotTo(gbytes.Say("any error message"))

			func() {
				defer func() {
					_ = recover()
				}()
				logger.Panic("panic message")
			}()
			logger.Panic("any panic message")
			Expect(mockWriter).NotTo(gbytes.Say("any panic message"))
		})
	})

	Describe("GetLogLevel", func() {
		It("should return current log level", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.SetLogLevel(usecase.LOG_DISABLED)
			Expect(logger.GetLogLevel()).To(Equal(usecase.LOG_DISABLED))

			logger.SetLogLevel(usecase.LOG_LEVEL_DEBUG)
			Expect(logger.GetLogLevel()).To(Equal(usecase.LOG_LEVEL_DEBUG))
		})
	})

	Describe("Panic", func() {
		It("should panic", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			Expect(func() {
				logger.Panic("panic message")
			}).To(Panic())
		})

		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			func() {
				defer func() {
					_ = recover()
				}()
				logger.Panic("panic message")
			}()

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the panic message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			func() {
				defer func() {
					_ = recover()
				}()
				logger.Panic("panic message")
			}()

			Expect(mockWriter).To(gbytes.Say("PNC panic message"))
		})
	})

	Describe("Panicf", func() {
		It("should panic", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			Expect(func() {
				logger.Panicf("%s panic message", "formatted")
			}).To(Panic())
		})

		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			func() {
				defer func() {
					_ = recover()
				}()
				logger.Panicf("%s panic message", "formatted")
			}()

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the formatted panic message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			func() {
				defer func() {
					_ = recover()
				}()
				logger.Panicf("%s panic message", "formatted")
			}()

			Expect(mockWriter).To(gbytes.Say("PNC formatted panic message"))
		})
	})

	Describe("Error", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Error("error message")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the error message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Error("error message")

			Expect(mockWriter).To(gbytes.Say("ERR error message"))
		})
	})

	Describe("Errorf", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Errorf("%s error message", "formatted")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the error message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Errorf("%s error message", "formatted")

			Expect(mockWriter).To(gbytes.Say("ERR formatted error message"))
		})
	})

	Describe("Info", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Info("info message")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the info message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Info("info message")

			Expect(mockWriter).To(gbytes.Say("INF info message"))
		})
	})

	Describe("Infof", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Infof("%s info message", "formatted")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the info message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Infof("%s info message", "formatted")

			Expect(mockWriter).To(gbytes.Say("INF formatted info message"))
		})
	})

	Describe("Debug", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Debug("debug message")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the debug message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Debug("debug message")

			Expect(mockWriter).To(gbytes.Say("DBG debug message"))
		})
	})

	Describe("Debugf", func() {
		It("should log with current time", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Debugf("%s debug message", "formatted")

			Expect(mockWriter).To(gbytes.Say(RFC3339_TIME_REGEX))
		})

		It("should log the debug message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			logger.Debugf("%s debug message", "formatted")

			Expect(mockWriter).To(gbytes.Say("DBG formatted debug message"))
		})
	})

	Describe("WithInterface", func() {
		It("should append key value pair to the log message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithInterface("anyKey", "any value")

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anyKey=\"any value\""))
		})

		It("should return a reusable logger with key-value context", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithInterface("anyKey", "any value")

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anyKey=\"any value\""))

			contextedLogger.Infof("any %s info message", "formatted")
			Expect(mockWriter).To(gbytes.Say("INF any formatted info message anyKey=\"any value\""))
		})

		It("should preserve the original logger context", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithInterface("anyKey", "any value")

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anyKey=\"any value\""))

			logger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message"))
		})

		It("should be chainable", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithInterface("anyKey", "any value").WithInterface("anotherKey", "another value")

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anotherKey=\"another value\" anyKey=\"any value\""))
		})
	})

	Describe("WithFields", func() {
		It("should append fields to the log message", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithFields(usecase.LogFields{
				"anyKey":     "any value",
				"anotherKey": "another value",
			})

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anotherKey=\"another value\" anyKey=\"any value\""))
		})

		It("should return a reusable logger with fields context", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithFields(usecase.LogFields{
				"anyKey":     "any value",
				"anotherKey": "another value",
			})

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anotherKey=\"another value\" anyKey=\"any value\""))

			contextedLogger.Infof("any %s info message", "formatted")
			Expect(mockWriter).To(gbytes.Say("INF any formatted info message anotherKey=\"another value\" anyKey=\"any value\""))
		})

		It("should preserve the original logger context", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithFields(usecase.LogFields{
				"anyKey":     "any value",
				"anotherKey": "another value",
			})

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anotherKey=\"another value\" anyKey=\"any value\""))

			logger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message"))
		})

		It("should be chainable", func() {
			mockWriter := NewMockWriter()
			logger := infrastructure.NewZerologLogger(mockWriter)

			contextedLogger := logger.WithFields(usecase.LogFields{
				"anyKey":     "any value",
				"anotherKey": "another value",
			}).WithFields((usecase.LogFields{
				"moreKey": "more value",
			}))

			contextedLogger.Info("any info message")
			Expect(mockWriter).To(gbytes.Say("INF any info message anotherKey=\"another value\" anyKey=\"any value\" moreKey=\"more value\""))
		})
	})
})

func NewMockWriter() io.WriteCloser {
	return gbytes.NewBuffer()
}
