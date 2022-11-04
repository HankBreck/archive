/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { ContactInfo } from "../contractregistry/contract";

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

const baseMsgRegisterContract: object = {
  creator: "",
  description: "",
  authors: "",
  moreInfoUri: "",
  templateUri: "",
  templateSchemaUri: "",
};

export const MsgRegisterContract = {
  encode(
    message: MsgRegisterContract,
    writer: Writer = Writer.create()
  ): Writer {
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
      ContactInfo.encode(
        message.contactInfo,
        writer.uint32(34).fork()
      ).ldelim();
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

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterContract } as MsgRegisterContract;
    message.authors = [];
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
    const message = { ...baseMsgRegisterContract } as MsgRegisterContract;
    message.authors = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.authors !== undefined && object.authors !== null) {
      for (const e of object.authors) {
        message.authors.push(String(e));
      }
    }
    if (object.contactInfo !== undefined && object.contactInfo !== null) {
      message.contactInfo = ContactInfo.fromJSON(object.contactInfo);
    } else {
      message.contactInfo = undefined;
    }
    if (object.moreInfoUri !== undefined && object.moreInfoUri !== null) {
      message.moreInfoUri = String(object.moreInfoUri);
    } else {
      message.moreInfoUri = "";
    }
    if (
      object.signingDataSchema !== undefined &&
      object.signingDataSchema !== null
    ) {
      message.signingDataSchema = bytesFromBase64(object.signingDataSchema);
    }
    if (object.templateUri !== undefined && object.templateUri !== null) {
      message.templateUri = String(object.templateUri);
    } else {
      message.templateUri = "";
    }
    if (
      object.templateSchemaUri !== undefined &&
      object.templateSchemaUri !== null
    ) {
      message.templateSchemaUri = String(object.templateSchemaUri);
    } else {
      message.templateSchemaUri = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.description !== undefined &&
      (obj.description = message.description);
    if (message.authors) {
      obj.authors = message.authors.map((e) => e);
    } else {
      obj.authors = [];
    }
    message.contactInfo !== undefined &&
      (obj.contactInfo = message.contactInfo
        ? ContactInfo.toJSON(message.contactInfo)
        : undefined);
    message.moreInfoUri !== undefined &&
      (obj.moreInfoUri = message.moreInfoUri);
    message.signingDataSchema !== undefined &&
      (obj.signingDataSchema = base64FromBytes(
        message.signingDataSchema !== undefined
          ? message.signingDataSchema
          : new Uint8Array()
      ));
    message.templateUri !== undefined &&
      (obj.templateUri = message.templateUri);
    message.templateSchemaUri !== undefined &&
      (obj.templateSchemaUri = message.templateSchemaUri);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegisterContract>): MsgRegisterContract {
    const message = { ...baseMsgRegisterContract } as MsgRegisterContract;
    message.authors = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.authors !== undefined && object.authors !== null) {
      for (const e of object.authors) {
        message.authors.push(e);
      }
    }
    if (object.contactInfo !== undefined && object.contactInfo !== null) {
      message.contactInfo = ContactInfo.fromPartial(object.contactInfo);
    } else {
      message.contactInfo = undefined;
    }
    if (object.moreInfoUri !== undefined && object.moreInfoUri !== null) {
      message.moreInfoUri = object.moreInfoUri;
    } else {
      message.moreInfoUri = "";
    }
    if (
      object.signingDataSchema !== undefined &&
      object.signingDataSchema !== null
    ) {
      message.signingDataSchema = object.signingDataSchema;
    } else {
      message.signingDataSchema = new Uint8Array();
    }
    if (object.templateUri !== undefined && object.templateUri !== null) {
      message.templateUri = object.templateUri;
    } else {
      message.templateUri = "";
    }
    if (
      object.templateSchemaUri !== undefined &&
      object.templateSchemaUri !== null
    ) {
      message.templateSchemaUri = object.templateSchemaUri;
    } else {
      message.templateSchemaUri = "";
    }
    return message;
  },
};

const baseMsgRegisterContractResponse: object = { id: 0 };

export const MsgRegisterContractResponse = {
  encode(
    message: MsgRegisterContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterContractResponse,
    } as MsgRegisterContractResponse;
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
    const message = {
      ...baseMsgRegisterContractResponse,
    } as MsgRegisterContractResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgRegisterContractResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterContractResponse>
  ): MsgRegisterContractResponse {
    const message = {
      ...baseMsgRegisterContractResponse,
    } as MsgRegisterContractResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RegisterContract(
    request: MsgRegisterContract
  ): Promise<MsgRegisterContractResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RegisterContract(
    request: MsgRegisterContract
  ): Promise<MsgRegisterContractResponse> {
    const data = MsgRegisterContract.encode(request).finish();
    const promise = this.rpc.request(
      "archive.contractregistry.Msg",
      "RegisterContract",
      data
    );
    return promise.then((data) =>
      MsgRegisterContractResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
