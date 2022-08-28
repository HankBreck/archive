/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Ownership } from "../cda/cda";

export const protobufPackage = "archive.cda";

/**
 * message Ownership {
 *     string owner = 1;
 *     uint64 ownership = 2;
 *   }
 */
export interface MsgCreateCDA {
  creator: string;
  cid: string;
  ownership: Ownership[];
  expiration: number;
}

export interface MsgCreateCDAResponse {
  id: number;
}

const baseMsgCreateCDA: object = { creator: "", cid: "", expiration: 0 };

export const MsgCreateCDA = {
  encode(message: MsgCreateCDA, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    for (const v of message.ownership) {
      Ownership.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.expiration !== 0) {
      writer.uint32(32).uint64(message.expiration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.ownership.push(Ownership.decode(reader, reader.uint32()));
          break;
        case 4:
          message.expiration = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCDA {
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromJSON(e));
      }
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = Number(object.expiration);
    } else {
      message.expiration = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateCDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    if (message.ownership) {
      obj.ownership = message.ownership.map((e) =>
        e ? Ownership.toJSON(e) : undefined
      );
    } else {
      obj.ownership = [];
    }
    message.expiration !== undefined && (obj.expiration = message.expiration);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCDA>): MsgCreateCDA {
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    message.ownership = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      for (const e of object.ownership) {
        message.ownership.push(Ownership.fromPartial(e));
      }
    }
    if (object.expiration !== undefined && object.expiration !== null) {
      message.expiration = object.expiration;
    } else {
      message.expiration = 0;
    }
    return message;
  },
};

const baseMsgCreateCDAResponse: object = { id: 0 };

export const MsgCreateCDAResponse = {
  encode(
    message: MsgCreateCDAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
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

  fromJSON(object: any): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateCDAResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCDAResponse>): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
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
  CreateCDA(request: MsgCreateCDA): Promise<MsgCreateCDAResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateCDA(request: MsgCreateCDA): Promise<MsgCreateCDAResponse> {
    const data = MsgCreateCDA.encode(request).finish();
    const promise = this.rpc.request("archive.cda.Msg", "CreateCDA", data);
    return promise.then((data) =>
      MsgCreateCDAResponse.decode(new Reader(data))
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
