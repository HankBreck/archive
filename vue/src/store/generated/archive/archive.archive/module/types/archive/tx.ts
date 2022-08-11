/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "archive.archive";

export interface MsgCreateCDA {
  creator: string;
  cid: string;
}

export interface MsgCreateCDAResponse {}

const baseMsgCreateCDA: object = { creator: "", cid: "" };

export const MsgCreateCDA = {
  encode(message: MsgCreateCDA, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
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
    return message;
  },

  toJSON(message: MsgCreateCDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCDA>): MsgCreateCDA {
    const message = { ...baseMsgCreateCDA } as MsgCreateCDA;
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
    return message;
  },
};

const baseMsgCreateCDAResponse: object = {};

export const MsgCreateCDAResponse = {
  encode(_: MsgCreateCDAResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCDAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
    return message;
  },

  toJSON(_: MsgCreateCDAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateCDAResponse>): MsgCreateCDAResponse {
    const message = { ...baseMsgCreateCDAResponse } as MsgCreateCDAResponse;
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
    const promise = this.rpc.request("archive.archive.Msg", "CreateCDA", data);
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
