import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Translate, TranslateCreateData } from '../InkspanDocumentTypes';
declare class TranslateEntity extends InkspanDocumentEntityBase<Translate> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: TranslateEntity): TranslateEntity;
    create(this: any, reqdata?: TranslateCreateData, ctrl?: Control): Promise<TranslateEntity>;
}
export { TranslateEntity };
