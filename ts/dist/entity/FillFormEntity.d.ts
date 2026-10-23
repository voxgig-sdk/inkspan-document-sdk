import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { FillForm, FillFormCreateData } from '../InkspanDocumentTypes';
declare class FillFormEntity extends InkspanDocumentEntityBase<FillForm> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: FillFormEntity): FillFormEntity;
    create(this: any, reqdata?: FillFormCreateData, ctrl?: Control): Promise<FillFormEntity>;
}
export { FillFormEntity };
