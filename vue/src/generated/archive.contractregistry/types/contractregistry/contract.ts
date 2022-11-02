/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "archive.contractregistry";

export enum ContactMethod {
  /** Phone - People won't want to publish their phone number on here */
  Phone = 0,
  Email = 1,
  UNRECOGNIZED = -1,
}

export function contactMethodFromJSON(object: any): ContactMethod {
  switch (object) {
    case 0:
    case "Phone":
      return ContactMethod.Phone;
    case 1:
    case "Email":
      return ContactMethod.Email;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ContactMethod.UNRECOGNIZED;
  }
}

export function contactMethodToJSON(object: ContactMethod): string {
  switch (object) {
    case ContactMethod.Phone:
      return "Phone";
    case ContactMethod.Email:
      return "Email";
    default:
      return "UNKNOWN";
  }
}

export interface Contract {
  id: string;
  description: string;
  authors: string;
  contactInfo: ContactInfo | undefined;
  moreInfoUri: string;
  /** Almost certainly want this stored separate from the contract */
  signingDataSchema: Uint8Array;
  templateUri: string;
  templateSchemaUri: string;
  version: string;
}

export interface ContactInfo {
  method: ContactMethod;
  value: string;
}

const baseContract: object = {
  id: "",
  description: "",
  authors: "",
  moreInfoUri: "",
  templateUri: "",
  templateSchemaUri: "",
  version: "",
};

export const Contract = {
  encode(message: Contract, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.authors !== "") {
      writer.uint32(26).string(message.authors);
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
    if (message.version !== "") {
      writer.uint32(74).string(message.version);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Contract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContract } as Contract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.authors = reader.string();
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
        case 9:
          message.version = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Contract {
    const message = { ...baseContract } as Contract;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.authors !== undefined && object.authors !== null) {
      message.authors = String(object.authors);
    } else {
      message.authors = "";
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
    if (object.version !== undefined && object.version !== null) {
      message.version = String(object.version);
    } else {
      message.version = "";
    }
    return message;
  },

  toJSON(message: Contract): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.description !== undefined &&
      (obj.description = message.description);
    message.authors !== undefined && (obj.authors = message.authors);
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
    message.version !== undefined && (obj.version = message.version);
    return obj;
  },

  fromPartial(object: DeepPartial<Contract>): Contract {
    const message = { ...baseContract } as Contract;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.authors !== undefined && object.authors !== null) {
      message.authors = object.authors;
    } else {
      message.authors = "";
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
    if (object.version !== undefined && object.version !== null) {
      message.version = object.version;
    } else {
      message.version = "";
    }
    return message;
  },
};

const baseContactInfo: object = { method: 0, value: "" };

export const ContactInfo = {
  encode(message: ContactInfo, writer: Writer = Writer.create()): Writer {
    if (message.method !== 0) {
      writer.uint32(8).int32(message.method);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ContactInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContactInfo } as ContactInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.method = reader.int32() as any;
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContactInfo {
    const message = { ...baseContactInfo } as ContactInfo;
    if (object.method !== undefined && object.method !== null) {
      message.method = contactMethodFromJSON(object.method);
    } else {
      message.method = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ContactInfo): unknown {
    const obj: any = {};
    message.method !== undefined &&
      (obj.method = contactMethodToJSON(message.method));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<ContactInfo>): ContactInfo {
    const message = { ...baseContactInfo } as ContactInfo;
    if (object.method !== undefined && object.method !== null) {
      message.method = object.method;
    } else {
      message.method = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
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
