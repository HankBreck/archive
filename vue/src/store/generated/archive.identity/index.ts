import { Client, registry, MissingWalletError } from 'archive-client-ts'

import { HashEntry } from "archive-client-ts/archive.identity/types"
import { Certificate } from "archive-client-ts/archive.identity/types"
import { Issuer } from "archive-client-ts/archive.identity/types"
import { Params } from "archive-client-ts/archive.identity/types"


export { HashEntry, Certificate, Issuer, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				IdentityMembers: {},
				Issuers: {},
				IssuerInfo: {},
				Identity: {},
				Operators: {},
				
				_Structure: {
						HashEntry: getStructure(HashEntry.fromPartial({})),
						Certificate: getStructure(Certificate.fromPartial({})),
						Issuer: getStructure(Issuer.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getIdentityMembers: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.IdentityMembers[JSON.stringify(params)] ?? {}
		},
				getIssuers: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Issuers[JSON.stringify(params)] ?? {}
		},
				getIssuerInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.IssuerInfo[JSON.stringify(params)] ?? {}
		},
				getIdentity: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Identity[JSON.stringify(params)] ?? {}
		},
				getOperators: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Operators[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: archive.identity initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIdentityMembers({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryIdentityMembers( key.id,  key.is_pending, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ArchiveIdentity.query.queryIdentityMembers( key.id,  key.is_pending, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'IdentityMembers', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIdentityMembers', payload: { options: { all }, params: {...key},query }})
				return getters['getIdentityMembers']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIdentityMembers API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIssuers({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryIssuers(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ArchiveIdentity.query.queryIssuers({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Issuers', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIssuers', payload: { options: { all }, params: {...key},query }})
				return getters['getIssuers']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIssuers API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIssuerInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryIssuerInfo( key.issuer)).data
				
					
				commit('QUERY', { query: 'IssuerInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIssuerInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getIssuerInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIssuerInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryIdentity({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryIdentity( key.id)).data
				
					
				commit('QUERY', { query: 'Identity', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryIdentity', payload: { options: { all }, params: {...key},query }})
				return getters['getIdentity']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryIdentity API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOperators({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ArchiveIdentity.query.queryOperators( key.id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ArchiveIdentity.query.queryOperators( key.id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Operators', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOperators', payload: { options: { all }, params: {...key},query }})
				return getters['getOperators']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOperators API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateOperators({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgUpdateOperators({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOperators:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOperators:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegisterIssuer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgRegisterIssuer({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterIssuer:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegisterIssuer:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRenounceIdentity({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgRenounceIdentity({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRenounceIdentity:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRenounceIdentity:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRejectIdentity({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgRejectIdentity({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectIdentity:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRejectIdentity:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateMembers({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgUpdateMembers({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMembers:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateMembers:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgIssueCertificate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgIssueCertificate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgIssueCertificate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgIssueCertificate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAcceptIdentity({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgAcceptIdentity({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAcceptIdentity:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAcceptIdentity:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddIdentityMember({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgAddIdentityMember({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddIdentityMember:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddIdentityMember:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRevokeIdentity({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ArchiveIdentity.tx.sendMsgRevokeIdentity({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeIdentity:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevokeIdentity:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateOperators({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgUpdateOperators({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOperators:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOperators:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegisterIssuer({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgRegisterIssuer({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterIssuer:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegisterIssuer:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRenounceIdentity({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgRenounceIdentity({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRenounceIdentity:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRenounceIdentity:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRejectIdentity({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgRejectIdentity({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectIdentity:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRejectIdentity:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateMembers({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgUpdateMembers({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMembers:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateMembers:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgIssueCertificate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgIssueCertificate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgIssueCertificate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgIssueCertificate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAcceptIdentity({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgAcceptIdentity({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAcceptIdentity:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAcceptIdentity:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddIdentityMember({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgAddIdentityMember({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddIdentityMember:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddIdentityMember:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRevokeIdentity({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ArchiveIdentity.tx.msgRevokeIdentity({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeIdentity:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevokeIdentity:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
