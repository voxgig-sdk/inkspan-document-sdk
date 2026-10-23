import { InkspanDocumentEntityBase } from '../InkspanDocumentEntityBase';
import type { InkspanDocumentSDK } from '../InkspanDocumentSDK';
import type { Control } from '../types';
import type { Watermark, WatermarkCreateData } from '../InkspanDocumentTypes';
declare class WatermarkEntity extends InkspanDocumentEntityBase<Watermark> {
    constructor(client: InkspanDocumentSDK, entopts: any);
    make(this: WatermarkEntity): WatermarkEntity;
    create(this: any, reqdata?: WatermarkCreateData, ctrl?: Control): Promise<WatermarkEntity>;
}
export { WatermarkEntity };
