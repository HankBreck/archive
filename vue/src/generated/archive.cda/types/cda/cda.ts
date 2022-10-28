/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

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
    default:
      return "UNKNOWN";
  }
}

export interface Ownership {
  owner: string;
  ownership: number;
}

const baseCDA: object = {
  creator: "",
  id: 0,
  signingParties: "",
  contractId: 0,
  legalMetadataUri: "",
  status: 0,
};

export const CDA = {
  encode(message: CDA, writer: Writer = Writer.create()): Writer {
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
      Timestamp.encode(
        toTimestamp(message.utcExpireTime),
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(56).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CDA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCDA } as CDA;
    message.signingParties = [];
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
          message.utcExpireTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
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
    const message = { ...baseCDA } as CDA;
    message.signingParties = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.signingParties !== undefined && object.signingParties !== null) {
      for (const e of object.signingParties) {
        message.signingParties.push(String(e));
      }
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = Number(object.contractId);
    } else {
      message.contractId = 0;
    }
    if (
      object.legalMetadataUri !== undefined &&
      object.legalMetadataUri !== null
    ) {
      message.legalMetadataUri = String(object.legalMetadataUri);
    } else {
      message.legalMetadataUri = "";
    }
    if (object.utcExpireTime !== undefined && object.utcExpireTime !== null) {
      message.utcExpireTime = fromJsonTimestamp(object.utcExpireTime);
    } else {
      message.utcExpireTime = undefined;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = cDA_ContractStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: CDA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    if (message.signingParties) {
      obj.signingParties = message.signingParties.map((e) => e);
    } else {
      obj.signingParties = [];
    }
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.legalMetadataUri !== undefined &&
      (obj.legalMetadataUri = message.legalMetadataUri);
    message.utcExpireTime !== undefined &&
      (obj.utcExpireTime =
        message.utcExpireTime !== undefined
          ? message.utcExpireTime.toISOString()
          : null);
    message.status !== undefined &&
      (obj.status = cDA_ContractStatusToJSON(message.status));
    return obj;
  },

  fromPartial(object: DeepPartial<CDA>): CDA {
    const message = { ...baseCDA } as CDA;
    message.signingParties = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.signingParties !== undefined && object.signingParties !== null) {
      for (const e of object.signingParties) {
        message.signingParties.push(e);
      }
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = 0;
    }
    if (
      object.legalMetadataUri !== undefined &&
      object.legalMetadataUri !== null
    ) {
      message.legalMetadataUri = object.legalMetadataUri;
    } else {
      message.legalMetadataUri = "";
    }
    if (object.utcExpireTime !== undefined && object.utcExpireTime !== null) {
      message.utcExpireTime = object.utcExpireTime;
    } else {
      message.utcExpireTime = undefined;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

const baseOwnership: object = { owner: "", ownership: 0 };

export const Ownership = {
  encode(message: Ownership, writer: Writer = Writer.create()): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.ownership !== 0) {
      writer.uint32(16).uint64(message.ownership);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Ownership {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOwnership } as Ownership;
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
    const message = { ...baseOwnership } as Ownership;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      message.ownership = Number(object.ownership);
    } else {
      message.ownership = 0;
    }
    return message;
  },

  toJSON(message: Ownership): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.ownership !== undefined && (obj.ownership = message.ownership);
    return obj;
  },

  fromPartial(object: DeepPartial<Ownership>): Ownership {
    const message = { ...baseOwnership } as Ownership;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.ownership !== undefined && object.ownership !== null) {
      message.ownership = object.ownership;
    } else {
      message.ownership = 0;
    }
    return message;
  },
};

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

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
