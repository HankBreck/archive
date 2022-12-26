/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { ContactInfo } from "./contract";

export const protobufPackage = "archive.contractregistry";

export interface MsgRegisterContract {
  creator: string;
  description: string;
  authors: string[];
  contactInfo: ContactInfo | undefined;
  moreInfoUri: string;
  signingDataSchema: Uint8Array;
  templateUri: string;
  templateSchemaUri: string;
}

export interface MsgRegisterContractResponse {
  id: number;
}

function createBaseMsgRegisterContract(): MsgRegisterContract {
  return {
    creator: "",
    description: "",
    authors: [],
    contactInfo: undefined,
    moreInfoUri: "",
    signingDataSchema: new Uint8Array(),
    templateUri: "",
    templateSchemaUri: "",
  };
}

export const MsgRegisterContract = {
  encode(message: MsgRegisterContract, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    for (const v of message.authors) {
      writer.uint32(26).string(v!);
    }
    if (message.contactInfo !== undefined) {
      ContactInfo.encode(message.contactInfo, writer.uint32(34).fork()).ldelim();
    }
    if (message.moreInfoUri !== "") {
      writer.uint32(42).string(message.moreInfoUri);
    }
    if (message.signingDataSchema.length !== 0) {
      writer.uint32(50).bytes(message.signingDataSchema);
    }
    if (message.templateUri !== "") {
      writer.uint32(58).string(message.templateUri);
    }
    if (message.templateSchemaUri !== "") {
      writer.uint32(66).string(message.templateSchemaUri);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterContract {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterContract();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.authors.push(reader.string());
          break;
        case 4:
          message.contactInfo = ContactInfo.decode(reader, reader.uint32());
          break;
        case 5:
          message.moreInfoUri = reader.string();
          break;
        case 6:
          message.signingDataSchema = reader.bytes();
          break;
        case 7:
          message.templateUri = reader.string();
          break;
        case 8:
          message.templateSchemaUri = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterContract {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      description: isSet(object.description) ? String(object.description) : "",
      authors: Array.isArray(object?.authors) ? object.authors.map((e: any) => String(e)) : [],
      contactInfo: isSet(object.contactInfo) ? ContactInfo.fromJSON(object.contactInfo) : undefined,
      moreInfoUri: isSet(object.moreInfoUri) ? String(object.moreInfoUri) : "",
      signingDataSchema: isSet(object.signingDataSchema) ? bytesFromBase64(object.signingDataSchema) : new Uint8Array(),
      templateUri: isSet(object.templateUri) ? String(object.templateUri) : "",
      templateSchemaUri: isSet(object.templateSchemaUri) ? String(object.templateSchemaUri) : "",
    };
  },

  toJSON(message: MsgRegisterContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.description !== undefined && (obj.description = message.description);
    if (message.authors) {
      obj.authors = message.authors.map((e) => e);
    } else {
      obj.authors = [];
    }
    message.contactInfo !== undefined
      && (obj.contactInfo = message.contactInfo ? ContactInfo.toJSON(message.contactInfo) : undefined);
    message.moreInfoUri !== undefined && (obj.moreInfoUri = message.moreInfoUri);
    message.signingDataSchema !== undefined
      && (obj.signingDataSchema = base64FromBytes(
        message.signingDataSchema !== undefined ? message.signingDataSchema : new Uint8Array(),
      ));
    message.templateUri !== undefined && (obj.templateUri = message.templateUri);
    message.templateSchemaUri !== undefined && (obj.templateSchemaUri = message.templateSchemaUri);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterContract>, I>>(object: I): MsgRegisterContract {
    const message = createBaseMsgRegisterContract();
    message.creator = object.creator ?? "";
    message.description = object.description ?? "";
    message.authors = object.authors?.map((e) => e) || [];
    message.contactInfo = (object.contactInfo !== undefined && object.contactInfo !== null)
      ? ContactInfo.fromPartial(object.contactInfo)
      : undefined;
    message.moreInfoUri = object.moreInfoUri ?? "";
    message.signingDataSchema = object.signingDataSchema ?? new Uint8Array();
    message.templateUri = object.templateUri ?? "";
    message.templateSchemaUri = object.templateSchemaUri ?? "";
    return message;
  },
};

function createBaseMsgRegisterContractResponse(): MsgRegisterContractResponse {
  return { id: 0 };
}

export const MsgRegisterContractResponse = {
  encode(message: MsgRegisterContractResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterContractResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterContractResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterContractResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgRegisterContractResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterContractResponse>, I>>(object: I): MsgRegisterContractResponse {
    const message = createBaseMsgRegisterContractResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RegisterContract(request: MsgRegisterContract): Promise<MsgRegisterContractResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegisterContract = this.RegisterContract.bind(this);
  }
  RegisterContract(request: MsgRegisterContract): Promise<MsgRegisterContractResponse> {
    const data = MsgRegisterContract.encode(request).finish();
    const promise = this.rpc.request("archive.contractregistry.Msg", "RegisterContract", data);
    return promise.then((data) => MsgRegisterContractResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
