import axios from "axios";
import PurchasesFileProcessorException from "./exceptions/PurchaseFileProcessorException";

export default class PurchasesFileProcessor {
  private _httpClient: any;

  constructor(httpClient = axios) {
    this._httpClient = httpClient;
  }

  async process(formData: FormData) {
    try {
      const response = await this._httpClient.post("http://localhost:8080/process-file", formData);
      const data = response.data.aggregated;
      const purchases = Object.values(data).map((clientData: any) => clientData.Purchases);
      return {
        clientNames: Object.keys(data),
        purchases: purchases.flat()
      }
    } catch (errRequest) {
      throw new PurchasesFileProcessorException((errRequest as Error).message);
    }
  } 
}