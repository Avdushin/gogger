package lib

// This is a writer function
func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.fileWriter.Write(p)
}
