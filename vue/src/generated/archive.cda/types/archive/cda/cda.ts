/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "archive.cda";

export interface CDA {
  creator: string;
  id: number;
  signingParties: string[];
  contractId: number;
  legalMetadataUri: string;
  utcExpireTime: Date | undefined;
  status: CDA_ContractStatus;
}

export enum CDA_ContractStatus {
  Pending = 0,
  Finalized = 1,
  Voided = 2,
  UNRECOGNIZED = -1,
}

export function cDA_ContractStatusFromJSON(object: any): CDA_ContractStatus {
  switch (object) {
    case 0:
    case "Pending":
      return CDA_ContractStatus.Pending;
    case 1:
    case "Finalized":
      return CDA_ContractStatus.Finalized;
    case 2:
    case "Voided":
      return CDA_ContractStatus.Voided;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CDA_ContractStatus.UNRECOGNIZED;
  }
}

export function cDA_ContractStatusToJSON(object: CDA_ContractStatus): string {
  switch (object) {
    case CDA_ContractStatus.Pending:
      return "Pending";
    case CDA_ContractStatus.Finalized:
      return "Finalized";
    case CDA_ContractStatus.Voided:
      return "Voided";
    case CDA_ContractStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Ownership {
  owner: string;
  ownership: number;
}

function createBaseCDA(): CDA {
  return {
    creator: "",
    id: 0,
    signingParties: [],
    contractId: 0,
    legalMetadataUri: "",
    utcExpireTime: undefined,
    status: 0,
  };
}

export const CDA = {
  encode(message: CDA, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    for (const v of message.signingParties) {
      writer.uint32(26).string(v!);
    }
    if (message.contractId !== 0) {
      writer.uint32(32).uint64(message.contractId);
    }
    if (message.legalMetadataUri !== "") {
      writer.uint32(42).string(message.legalMetadataUri);
    }
    if (message.utcExpireTime !== undefined) {
      Timestamp.encode(toTimestamp(message.utcExpireTime), writer.uint32(50).fork()).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(56).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CDA {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCDA();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.signingParties.push(reader.string());
          break;
        case 4:
          message.contractId = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.legalMetadataUri = reader.string();
          break;
        case 6:
          message.utcExpireTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 7:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CDA {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      signingParties: Array.isArray(object?.signingParties) ? object.signingParties.map((e: any) => String(e)) : [],
      contractId: isSet(object.contractId) ? Number(object.contractId) : 0,
      legalMetadataUri: isSet(object.legalMetadataUri) ? String(object.legalMetadataUri) : "",
      utcExpireTime: isSet(object.utcExpireTime) ? fromJsonTimestamp(object.utcExpireTime) : undefined,
      status: isSet(object.status) ? cDA_ContractStatusFromJSON(object.status) : 0,
    };
  },

  toJSON(message: CDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    if (message.signingParties) {
      obj.signingParties = message.signingParties.map((e) => e);
    } else {
      obj.signingParties = [];
    }
    message.contractId !== undefined && (obj.contractId = Math.round(message.contractId));
    message.legalMetadataUri !== undefined && (obj.legalMetadataUri = message.legalMetadataUri);
    message.utcExpireTime !== undefined && (obj.utcExpireTime = message.utcExpireTime.toISOString());
    message.status !== undefined && (obj.status = cDA_ContractStatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CDA>, I>>(object: I): CDA {
    const message = createBaseCDA();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.signingParties = object.signingParties?.map((e) => e) || [];
    message.contractId = object.contractId ?? 0;
    message.legalMetadataUri = object.legalMetadataUri ?? "";
    message.utcExpireTime = object.utcExpireTime ?? undefined;
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseOwnership(): Ownership {
  return { owner: "", ownership: 0 };
}

export const Ownership = {
  encode(message: Ownership, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.ownership !== 0) {
      writer.uint32(16).uint64(message.ownership);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Ownership {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOwnership();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.ownership = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ownership {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      ownership: isSet(object.ownership) ? Number(object.ownership) : 0,
    };
  },

  toJSON(message: Ownership): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.ownership !== undefined && (obj.ownership = Math.round(message.ownership));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Ownership>, I>>(object: I): Ownership {
    const message = createBaseOwnership();
    message.owner = object.owner ?? "";
    message.ownership = object.ownership ?? 0;
    return message;
  },
};

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

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
