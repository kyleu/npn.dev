namespace auth {
  export interface Auth {
    readonly type: string;
    readonly config: object;
  }

  export interface Basic {
    readonly username: string;
    readonly password: string;
    readonly showPassword: boolean;
  }
}
