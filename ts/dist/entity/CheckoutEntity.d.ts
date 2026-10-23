import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Checkout, CheckoutCreateData } from '../InkspanDocumentTypes';
declare class CheckoutEntity extends InkspanDocumentEntityBase<Checkout> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: CheckoutEntity): CheckoutEntity;
    create(this: any, reqdata?: CheckoutCreateData, ctrl?: Control): Promise<CheckoutEntity>;
}
export { CheckoutEntity };
