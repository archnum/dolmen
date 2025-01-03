/*
####### dolmen (c) 2024 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package store

type (
	Config struct {
		Config     map[string]any `ms:"config"`
		Type       string         `ms:"type"`
		Namespaces []string       `ms:"namespaces"`
		Disabled   bool           `ms:"disabled"`
	}
)

/*
####### END ############################################################################################################
*/
