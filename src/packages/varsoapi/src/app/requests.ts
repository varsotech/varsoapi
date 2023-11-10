/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Album, Post } from "./base";

export const protobufPackage = "varso";

export interface GetPostsResponse {
  posts: Post[];
}

export interface CreatePostRequest {
  spotifyContentUuid: string;
}

export interface CatalogSearchRequest {
  query: string;
}

export interface CatalogSearchResponse {
  albums: Album[];
}

function createBaseGetPostsResponse(): GetPostsResponse {
  return { posts: [] };
}

export const GetPostsResponse = {
  encode(message: GetPostsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.posts) {
      Post.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPostsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPostsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.posts.push(Post.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPostsResponse {
    return { posts: Array.isArray(object?.posts) ? object.posts.map((e: any) => Post.fromJSON(e)) : [] };
  },

  toJSON(message: GetPostsResponse): unknown {
    const obj: any = {};
    if (message.posts?.length) {
      obj.posts = message.posts.map((e) => Post.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPostsResponse>, I>>(base?: I): GetPostsResponse {
    return GetPostsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetPostsResponse>, I>>(object: I): GetPostsResponse {
    const message = createBaseGetPostsResponse();
    message.posts = object.posts?.map((e) => Post.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreatePostRequest(): CreatePostRequest {
  return { spotifyContentUuid: "" };
}

export const CreatePostRequest = {
  encode(message: CreatePostRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.spotifyContentUuid !== "") {
      writer.uint32(10).string(message.spotifyContentUuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePostRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePostRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.spotifyContentUuid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreatePostRequest {
    return { spotifyContentUuid: isSet(object.spotifyContentUuid) ? String(object.spotifyContentUuid) : "" };
  },

  toJSON(message: CreatePostRequest): unknown {
    const obj: any = {};
    if (message.spotifyContentUuid !== "") {
      obj.spotifyContentUuid = message.spotifyContentUuid;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreatePostRequest>, I>>(base?: I): CreatePostRequest {
    return CreatePostRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreatePostRequest>, I>>(object: I): CreatePostRequest {
    const message = createBaseCreatePostRequest();
    message.spotifyContentUuid = object.spotifyContentUuid ?? "";
    return message;
  },
};

function createBaseCatalogSearchRequest(): CatalogSearchRequest {
  return { query: "" };
}

export const CatalogSearchRequest = {
  encode(message: CatalogSearchRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CatalogSearchRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCatalogSearchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.query = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CatalogSearchRequest {
    return { query: isSet(object.query) ? String(object.query) : "" };
  },

  toJSON(message: CatalogSearchRequest): unknown {
    const obj: any = {};
    if (message.query !== "") {
      obj.query = message.query;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CatalogSearchRequest>, I>>(base?: I): CatalogSearchRequest {
    return CatalogSearchRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CatalogSearchRequest>, I>>(object: I): CatalogSearchRequest {
    const message = createBaseCatalogSearchRequest();
    message.query = object.query ?? "";
    return message;
  },
};

function createBaseCatalogSearchResponse(): CatalogSearchResponse {
  return { albums: [] };
}

export const CatalogSearchResponse = {
  encode(message: CatalogSearchResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.albums) {
      Album.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CatalogSearchResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCatalogSearchResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.albums.push(Album.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CatalogSearchResponse {
    return { albums: Array.isArray(object?.albums) ? object.albums.map((e: any) => Album.fromJSON(e)) : [] };
  },

  toJSON(message: CatalogSearchResponse): unknown {
    const obj: any = {};
    if (message.albums?.length) {
      obj.albums = message.albums.map((e) => Album.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CatalogSearchResponse>, I>>(base?: I): CatalogSearchResponse {
    return CatalogSearchResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CatalogSearchResponse>, I>>(object: I): CatalogSearchResponse {
    const message = createBaseCatalogSearchResponse();
    message.albums = object.albums?.map((e) => Album.fromPartial(e)) || [];
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
