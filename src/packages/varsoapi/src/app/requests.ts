/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Organization, RSSFeed } from "./base";

export const protobufPackage = "varso";

export interface GetOrganizationsResponse {
  organizations: Organization[];
}

export interface GetNewsResponse {
  feeds: RSSFeed[];
}

function createBaseGetOrganizationsResponse(): GetOrganizationsResponse {
  return { organizations: [] };
}

export const GetOrganizationsResponse = {
  encode(message: GetOrganizationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.organizations) {
      Organization.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrganizationsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrganizationsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.organizations.push(Organization.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetOrganizationsResponse {
    return {
      organizations: Array.isArray(object?.organizations)
        ? object.organizations.map((e: any) => Organization.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetOrganizationsResponse): unknown {
    const obj: any = {};
    if (message.organizations?.length) {
      obj.organizations = message.organizations.map((e) => Organization.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(base?: I): GetOrganizationsResponse {
    return GetOrganizationsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(object: I): GetOrganizationsResponse {
    const message = createBaseGetOrganizationsResponse();
    message.organizations = object.organizations?.map((e) => Organization.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetNewsResponse(): GetNewsResponse {
  return { feeds: [] };
}

export const GetNewsResponse = {
  encode(message: GetNewsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.feeds) {
      RSSFeed.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.feeds.push(RSSFeed.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetNewsResponse {
    return { feeds: Array.isArray(object?.feeds) ? object.feeds.map((e: any) => RSSFeed.fromJSON(e)) : [] };
  },

  toJSON(message: GetNewsResponse): unknown {
    const obj: any = {};
    if (message.feeds?.length) {
      obj.feeds = message.feeds.map((e) => RSSFeed.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetNewsResponse>, I>>(base?: I): GetNewsResponse {
    return GetNewsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetNewsResponse>, I>>(object: I): GetNewsResponse {
    const message = createBaseGetNewsResponse();
    message.feeds = object.feeds?.map((e) => RSSFeed.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };
