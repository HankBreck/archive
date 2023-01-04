// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgAddIdentityMember } from "./types/archive/identity/tx";
import { MsgUpdateMembers } from "./types/archive/identity/tx";
import { MsgUpdateOperators } from "./types/archive/identity/tx";
import { MsgRejectIdentity } from "./types/archive/identity/tx";
import { MsgRevokeIdentity } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";


export { MsgIssueCertificate, MsgRenounceIdentity, MsgAddIdentityMember, MsgUpdateMembers, MsgUpdateOperators, MsgRejectIdentity, MsgRevokeIdentity, MsgRegisterIssuer, MsgAcceptIdentity };

type sendMsgIssueCertificateParams = {
  value: MsgIssueCertificate,
  fee?: StdFee,
  memo?: string
};

type sendMsgRenounceIdentityParams = {
  value: MsgRenounceIdentity,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddIdentityMemberParams = {
  value: MsgAddIdentityMember,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateMembersParams = {
  value: MsgUpdateMembers,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateOperatorsParams = {
  value: MsgUpdateOperators,
  fee?: StdFee,
  memo?: string
};

type sendMsgRejectIdentityParams = {
  value: MsgRejectIdentity,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeIdentityParams = {
  value: MsgRevokeIdentity,
  fee?: StdFee,
  memo?: string
};

type sendMsgRegisterIssuerParams = {
  value: MsgRegisterIssuer,
  fee?: StdFee,
  memo?: string
};

type sendMsgAcceptIdentityParams = {
  value: MsgAcceptIdentity,
  fee?: StdFee,
  memo?: string
};


type msgIssueCertificateParams = {
  value: MsgIssueCertificate,
};

type msgRenounceIdentityParams = {
  value: MsgRenounceIdentity,
};

type msgAddIdentityMemberParams = {
  value: MsgAddIdentityMember,
};

type msgUpdateMembersParams = {
  value: MsgUpdateMembers,
};

type msgUpdateOperatorsParams = {
  value: MsgUpdateOperators,
};

type msgRejectIdentityParams = {
  value: MsgRejectIdentity,
};

type msgRevokeIdentityParams = {
  value: MsgRevokeIdentity,
};

type msgRegisterIssuerParams = {
  value: MsgRegisterIssuer,
};

type msgAcceptIdentityParams = {
  value: MsgAcceptIdentity,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgIssueCertificate({ value, fee, memo }: sendMsgIssueCertificateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgIssueCertificate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgIssueCertificate({ value: MsgIssueCertificate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgIssueCertificate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRenounceIdentity({ value, fee, memo }: sendMsgRenounceIdentityParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRenounceIdentity: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRenounceIdentity({ value: MsgRenounceIdentity.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRenounceIdentity: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddIdentityMember({ value, fee, memo }: sendMsgAddIdentityMemberParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddIdentityMember: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddIdentityMember({ value: MsgAddIdentityMember.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddIdentityMember: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateMembers({ value, fee, memo }: sendMsgUpdateMembersParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateMembers: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateMembers({ value: MsgUpdateMembers.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateMembers: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateOperators({ value, fee, memo }: sendMsgUpdateOperatorsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateOperators: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateOperators({ value: MsgUpdateOperators.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateOperators: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRejectIdentity({ value, fee, memo }: sendMsgRejectIdentityParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRejectIdentity: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRejectIdentity({ value: MsgRejectIdentity.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRejectIdentity: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeIdentity({ value, fee, memo }: sendMsgRevokeIdentityParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeIdentity: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeIdentity({ value: MsgRevokeIdentity.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeIdentity: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRegisterIssuer({ value, fee, memo }: sendMsgRegisterIssuerParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRegisterIssuer: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRegisterIssuer({ value: MsgRegisterIssuer.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRegisterIssuer: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAcceptIdentity({ value, fee, memo }: sendMsgAcceptIdentityParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAcceptIdentity: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAcceptIdentity({ value: MsgAcceptIdentity.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAcceptIdentity: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgIssueCertificate({ value }: msgIssueCertificateParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgIssueCertificate", value: MsgIssueCertificate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgIssueCertificate: Could not create message: ' + e.message)
			}
		},
		
		msgRenounceIdentity({ value }: msgRenounceIdentityParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgRenounceIdentity", value: MsgRenounceIdentity.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRenounceIdentity: Could not create message: ' + e.message)
			}
		},
		
		msgAddIdentityMember({ value }: msgAddIdentityMemberParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgAddIdentityMember", value: MsgAddIdentityMember.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddIdentityMember: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateMembers({ value }: msgUpdateMembersParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgUpdateMembers", value: MsgUpdateMembers.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateMembers: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateOperators({ value }: msgUpdateOperatorsParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgUpdateOperators", value: MsgUpdateOperators.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateOperators: Could not create message: ' + e.message)
			}
		},
		
		msgRejectIdentity({ value }: msgRejectIdentityParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgRejectIdentity", value: MsgRejectIdentity.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRejectIdentity: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeIdentity({ value }: msgRevokeIdentityParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgRevokeIdentity", value: MsgRevokeIdentity.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeIdentity: Could not create message: ' + e.message)
			}
		},
		
		msgRegisterIssuer({ value }: msgRegisterIssuerParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgRegisterIssuer", value: MsgRegisterIssuer.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRegisterIssuer: Could not create message: ' + e.message)
			}
		},
		
		msgAcceptIdentity({ value }: msgAcceptIdentityParams): EncodeObject {
			try {
				return { typeUrl: "/archive.identity.MsgAcceptIdentity", value: MsgAcceptIdentity.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAcceptIdentity: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			ArchiveIdentity: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;