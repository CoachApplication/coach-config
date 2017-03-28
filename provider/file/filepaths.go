package file

/**
 * FilePaths a FileConnector related abstraction that does the effort of mapping config key and scope value to actual
 * file paths.  This means that various file arrangements can easily be handled without having to rewrite the file
 * loading functionality
 */
type FilePaths interface {
	// Keys return a string list of available keys
	Keys() []string
	// Scopes return a string list of available scopes
	Scopes() []string
	// Find the path for a key-scope pair, if they are valid
	Path(key, scope string) (string, error)
}
