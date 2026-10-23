import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Signup, SignupCreateData } from '../InkspanDocumentTypes';
declare class SignupEntity extends InkspanDocumentEntityBase<Signup> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: SignupEntity): SignupEntity;
    create(this: any, reqdata?: SignupCreateData, ctrl?: Control): Promise<SignupEntity>;
}
export { SignupEntity };
