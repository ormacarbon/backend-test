export interface ResponsePattern { 
    httpCode: number; 
    data: any; 
    status: boolean; 
    error: unknown | string | null;
}