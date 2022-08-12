import { txClient, queryClient, MissingWalletError , registry} from './module'

import { CDA } from "./module/types/cda/cda"
import { Params } from "./module/types/cda/params"


export { CDA, Params };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				Cda: {},
				
				_Structure: {
						CDA: getStructure(CDA.fromPartial({})),
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
				getCda: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Cda[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: archive.cda initialized!')
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCda({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCda( key.id)).data
				
					
				commit('QUERY', { query: 'Cda', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCda', payload: { options: { all }, params: {...key},query }})
				return getters['getCda']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCda API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateCDA({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateCDA(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCDA:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCDA:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateCDA({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateCDA(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCDA:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCDA:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
