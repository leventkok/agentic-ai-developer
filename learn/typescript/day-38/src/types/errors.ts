export type NetworkError = {
    kind: "network";
    message: string;
    statusCode?: number;
  };
  
  export type ValidationError = {
    kind: "validation";
    field: string;
    message: string;
  };
  
  export type NotFoundError = {
    kind: "not_found";
    resource: string;
    id: string;
  };
  
  export type ApiError = NetworkError | ValidationError | NotFoundError;