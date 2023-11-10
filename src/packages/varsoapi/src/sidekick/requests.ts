/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Task } from "./base";

export const protobufPackage = "varso";

export interface GetTasksRequest {
}

export interface GetTasksResponse {
  tasks: Task[];
}

export interface CreateTaskRequest {
  title: string;
}

export interface CreateTaskResponse {
}

export interface UpdateTaskRequest {
  uuid: string;
  title: string;
}

export interface UpdateTaskResponse {
}

export interface DeleteTaskRequest {
  uuid: string;
}

export interface DeleteTaskResponse {
}

function createBaseGetTasksRequest(): GetTasksRequest {
  return {};
}

export const GetTasksRequest = {
  encode(_: GetTasksRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTasksRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTasksRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetTasksRequest {
    return {};
  },

  toJSON(_: GetTasksRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTasksRequest>, I>>(base?: I): GetTasksRequest {
    return GetTasksRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTasksRequest>, I>>(_: I): GetTasksRequest {
    const message = createBaseGetTasksRequest();
    return message;
  },
};

function createBaseGetTasksResponse(): GetTasksResponse {
  return { tasks: [] };
}

export const GetTasksResponse = {
  encode(message: GetTasksResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tasks) {
      Task.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTasksResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTasksResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tasks.push(Task.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetTasksResponse {
    return { tasks: Array.isArray(object?.tasks) ? object.tasks.map((e: any) => Task.fromJSON(e)) : [] };
  },

  toJSON(message: GetTasksResponse): unknown {
    const obj: any = {};
    if (message.tasks?.length) {
      obj.tasks = message.tasks.map((e) => Task.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetTasksResponse>, I>>(base?: I): GetTasksResponse {
    return GetTasksResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetTasksResponse>, I>>(object: I): GetTasksResponse {
    const message = createBaseGetTasksResponse();
    message.tasks = object.tasks?.map((e) => Task.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateTaskRequest(): CreateTaskRequest {
  return { title: "" };
}

export const CreateTaskRequest = {
  encode(message: CreateTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTaskRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateTaskRequest {
    return { title: isSet(object.title) ? String(object.title) : "" };
  },

  toJSON(message: CreateTaskRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTaskRequest>, I>>(base?: I): CreateTaskRequest {
    return CreateTaskRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateTaskRequest>, I>>(object: I): CreateTaskRequest {
    const message = createBaseCreateTaskRequest();
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseCreateTaskResponse(): CreateTaskResponse {
  return {};
}

export const CreateTaskResponse = {
  encode(_: CreateTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTaskResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTaskResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CreateTaskResponse {
    return {};
  },

  toJSON(_: CreateTaskResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateTaskResponse>, I>>(base?: I): CreateTaskResponse {
    return CreateTaskResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateTaskResponse>, I>>(_: I): CreateTaskResponse {
    const message = createBaseCreateTaskResponse();
    return message;
  },
};

function createBaseUpdateTaskRequest(): UpdateTaskRequest {
  return { uuid: "", title: "" };
}

export const UpdateTaskRequest = {
  encode(message: UpdateTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTaskRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uuid = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateTaskRequest {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      title: isSet(object.title) ? String(object.title) : "",
    };
  },

  toJSON(message: UpdateTaskRequest): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTaskRequest>, I>>(base?: I): UpdateTaskRequest {
    return UpdateTaskRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateTaskRequest>, I>>(object: I): UpdateTaskRequest {
    const message = createBaseUpdateTaskRequest();
    message.uuid = object.uuid ?? "";
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseUpdateTaskResponse(): UpdateTaskResponse {
  return {};
}

export const UpdateTaskResponse = {
  encode(_: UpdateTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTaskResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTaskResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): UpdateTaskResponse {
    return {};
  },

  toJSON(_: UpdateTaskResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateTaskResponse>, I>>(base?: I): UpdateTaskResponse {
    return UpdateTaskResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateTaskResponse>, I>>(_: I): UpdateTaskResponse {
    const message = createBaseUpdateTaskResponse();
    return message;
  },
};

function createBaseDeleteTaskRequest(): DeleteTaskRequest {
  return { uuid: "" };
}

export const DeleteTaskRequest = {
  encode(message: DeleteTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTaskRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uuid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteTaskRequest {
    return { uuid: isSet(object.uuid) ? String(object.uuid) : "" };
  },

  toJSON(message: DeleteTaskRequest): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTaskRequest>, I>>(base?: I): DeleteTaskRequest {
    return DeleteTaskRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteTaskRequest>, I>>(object: I): DeleteTaskRequest {
    const message = createBaseDeleteTaskRequest();
    message.uuid = object.uuid ?? "";
    return message;
  },
};

function createBaseDeleteTaskResponse(): DeleteTaskResponse {
  return {};
}

export const DeleteTaskResponse = {
  encode(_: DeleteTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTaskResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTaskResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteTaskResponse {
    return {};
  },

  toJSON(_: DeleteTaskResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteTaskResponse>, I>>(base?: I): DeleteTaskResponse {
    return DeleteTaskResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteTaskResponse>, I>>(_: I): DeleteTaskResponse {
    const message = createBaseDeleteTaskResponse();
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
