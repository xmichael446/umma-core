import { Client, registry, MissingWalletError } from 'UmmadChain-umma-client-ts'

import { Denom } from "UmmadChain-umma-client-ts/ummadchain.umma.factory/types"
import { Params } from "UmmadChain-umma-client-ts/ummadchain.umma.factory/types"


export { Denom, Params };

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
				Denom: {},
				DenomAll: {},
				
				_Structure: {
						Denom: getStructure(Denom.fromPartial({})),
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
				getDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Denom[JSON.stringify(params)] ?? {}
		},
				getDenomAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DenomAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: ummadchain.umma.factory initialized!')
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
				let value= (await client.UmmadchainUmmaFactory.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.UmmadchainUmmaFactory.query.queryDenom( key.denom)).data
				
					
				commit('QUERY', { query: 'Denom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDenom API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDenomAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.UmmadchainUmmaFactory.query.queryDenomAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.UmmadchainUmmaFactory.query.queryDenomAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DenomAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDenomAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDenomAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDenomAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateOwner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.UmmadchainUmmaFactory.tx.sendMsgUpdateOwner({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOwner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOwner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateDenom({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.UmmadchainUmmaFactory.tx.sendMsgCreateDenom({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDenom:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateDenom:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgMintAndSendTokens({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.UmmadchainUmmaFactory.tx.sendMsgMintAndSendTokens({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintAndSendTokens:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMintAndSendTokens:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateDenom({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.UmmadchainUmmaFactory.tx.sendMsgUpdateDenom({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDenom:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateDenom:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBurnTokens({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.UmmadchainUmmaFactory.tx.sendMsgBurnTokens({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurnTokens:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBurnTokens:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateOwner({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.UmmadchainUmmaFactory.tx.msgUpdateOwner({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOwner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOwner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateDenom({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.UmmadchainUmmaFactory.tx.msgCreateDenom({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateDenom:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateDenom:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgMintAndSendTokens({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.UmmadchainUmmaFactory.tx.msgMintAndSendTokens({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintAndSendTokens:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMintAndSendTokens:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateDenom({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.UmmadchainUmmaFactory.tx.msgUpdateDenom({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDenom:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateDenom:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBurnTokens({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.UmmadchainUmmaFactory.tx.msgBurnTokens({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurnTokens:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBurnTokens:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
