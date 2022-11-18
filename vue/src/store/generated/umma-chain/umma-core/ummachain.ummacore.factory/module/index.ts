// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

// @ts-ignore
import { StdFee } from "@cosmjs/launchpad";
// @ts-ignore
import { SigningStargateClient } from "@cosmjs/stargate";
// @ts-ignore
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateDenom } from "./types/factory/tx";
import { MsgUpdateOwner } from "./types/factory/tx";
import { MsgBurnTokens } from "./types/factory/tx";
import { MsgCreateDenom } from "./types/factory/tx";
import { MsgMintAndSendTokens } from "./types/factory/tx";


const types = [
  ["/ummachain.ummacore.factory.MsgUpdateDenom", MsgUpdateDenom],
  ["/ummachain.ummacore.factory.MsgUpdateOwner", MsgUpdateOwner],
  ["/ummachain.ummacore.factory.MsgBurnTokens", MsgBurnTokens],
  ["/ummachain.ummacore.factory.MsgCreateDenom", MsgCreateDenom],
  ["/ummachain.ummacore.factory.MsgMintAndSendTokens", MsgMintAndSendTokens],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateDenom: (data: MsgUpdateDenom): EncodeObject => ({ typeUrl: "/ummachain.ummacore.factory.MsgUpdateDenom", value: MsgUpdateDenom.fromPartial( data ) }),
    msgUpdateOwner: (data: MsgUpdateOwner): EncodeObject => ({ typeUrl: "/ummachain.ummacore.factory.MsgUpdateOwner", value: MsgUpdateOwner.fromPartial( data ) }),
    msgBurnTokens: (data: MsgBurnTokens): EncodeObject => ({ typeUrl: "/ummachain.ummacore.factory.MsgBurnTokens", value: MsgBurnTokens.fromPartial( data ) }),
    msgCreateDenom: (data: MsgCreateDenom): EncodeObject => ({ typeUrl: "/ummachain.ummacore.factory.MsgCreateDenom", value: MsgCreateDenom.fromPartial( data ) }),
    msgMintAndSendTokens: (data: MsgMintAndSendTokens): EncodeObject => ({ typeUrl: "/ummachain.ummacore.factory.MsgMintAndSendTokens", value: MsgMintAndSendTokens.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
