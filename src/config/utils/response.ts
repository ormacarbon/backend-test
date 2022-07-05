export default class Response { 
    successfulRequest<DataType, ErrorType>(httpCode: number, data: DataType, error: ErrorType) {
        return {
            httpCode,
            data,
            status: true,
            error,
        };
    }

    badRequest<DataType, ErrorType>(httpCode: number = 400, data: DataType, error: ErrorType) {
        return {
            httpCode,
            data,
            status: false,
            error,
        };
    }

    error<ErrorType>(error: ErrorType) {
        console.log(error);

        return {
            httpCode: 500,
            data: null,
            status: false,
            error,
        }        
    }

}