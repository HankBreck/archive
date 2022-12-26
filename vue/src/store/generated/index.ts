// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ArchiveCda from './archive.cda'
import ArchiveContractregistry from './archive.contractregistry'
import ArchiveIdentity from './archive.identity'


export default { 
  ArchiveCda: load(ArchiveCda, 'archive.cda'),
  ArchiveContractregistry: load(ArchiveContractregistry, 'archive.contractregistry'),
  ArchiveIdentity: load(ArchiveIdentity, 'archive.identity'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}