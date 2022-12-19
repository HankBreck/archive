// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ArchiveIdentity from './archive.identity'
import ArchiveCda from './archive.cda'
import ArchiveContractregistry from './archive.contractregistry'


export default { 
  ArchiveIdentity: load(ArchiveIdentity, 'archive.identity'),
  ArchiveCda: load(ArchiveCda, 'archive.cda'),
  ArchiveContractregistry: load(ArchiveContractregistry, 'archive.contractregistry'),
  
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