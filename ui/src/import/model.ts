export interface ImportConfigFile {
  filename: string;
  size: number;
  contentType: string;
}

export interface ImportConfig {
  status: string;
  files: ImportConfigFile[]
}

export interface ImportResultFile {
  filename: string;
  type: string;
  value: any;
  error: string
}

export interface ImportResult {
  key: string;
  cfg?: ImportConfig
  results?: ImportResultFile[]
}
