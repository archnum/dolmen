/*
####### dolmen (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package mongodb

type (
	implStore struct {
		name string
	}
)

func New(name string, cfg map[string]any) (*implStore, error) {
	impl := &implStore{
		name: name,
	}

	return impl, nil
}

func (impl *implStore) Close() error {
	return nil
}

/*
####### END ############################################################################################################
*/
