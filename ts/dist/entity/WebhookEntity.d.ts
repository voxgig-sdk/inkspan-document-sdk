import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Webhook, WebhookCreateData } from '../InkspanDocumentTypes';
declare class WebhookEntity extends InkspanDocumentEntityBase<Webhook> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: WebhookEntity): WebhookEntity;
    create(this: any, reqdata?: WebhookCreateData, ctrl?: Control): Promise<WebhookEntity>;
}
export { WebhookEntity };
