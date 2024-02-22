import { ErrorHandler, Injectable } from '@angular/core'

@Injectable()
export class ErrorLogger implements ErrorHandler {
    handleError(error: Error): void {
        console.log(error);
        
    }
}
