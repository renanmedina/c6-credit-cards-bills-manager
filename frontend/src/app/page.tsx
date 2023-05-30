"use client"

import PurchasesFileProcessor from '@/analyzer/PurchasesFileProcessor';
import Image from 'next/image'
import { useState } from 'react'

export default function Home() {

  const [isProcessing, setIsProcessing] = useState(false);
  const [purchasesData, setPurchasesData] = useState<Array<any>>([]);
  const [clientNames, setClientNames] = useState<Array<string>>([]);
  const [selectedClient, setSelectedClient] = useState("");

  const UploadFile = async (fileInput: any) => {
    if (isProcessing) {
      return;
    }

    if (fileInput.value) {
      const form = document.getElementById("bill-upload-form");
      try {
        setIsProcessing(true);
        const processor = new PurchasesFileProcessor();
        const formData = new FormData(form as HTMLFormElement);
        const processedData = await processor.process(formData);
        setPurchasesData(processedData.purchases);
        setClientNames(processedData.clientNames);
      } catch (errProcessing) {
        alert(errProcessing);
      } finally {
        setIsProcessing(false);
      }
    }
  };

  const filterPurchases = (event: any) => {
    const selectedClient = event.target.value;
    setSelectedClient(selectedClient);
  }

  const renderItems = () => {
    if (isProcessing) { 
      return (
        <tr>
          <td 
            colSpan={5} 
            className='text-center border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>
            Carregando, aguarde ....
          </td>
        </tr>
      );
    }

    if (!purchasesData.length) {
      return (
        <tr>
          <td colSpan={5} className='text-center border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>Nenhuma compra encontrada</td>
        </tr>
      );
    }

    let filteredItems = purchasesData;
    if (selectedClient != "") {
      filteredItems = purchasesData.filter((item) => item.ClientName == selectedClient);
    }

    return filteredItems.map((purchase : any) => {
      const amountFormatted = Intl.NumberFormat("pt-BR", {style: "currency", currency: "BRL"}).format(purchase.Amount);
      return (
        <tr key={`${purchase.Description}-${purchase.Date}`}>
          <td className='border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>{ purchase.Date }</td>
          <td className='border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>{ purchase.ClientName }</td>
          <td className='border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>{ purchase.Description }</td>
          <td className='border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>{ purchase.Category }</td>
          <td className='border border-slate-300 dark:border-slate-700 p-2 text-slate-500 dark:text-slate-400'>{ amountFormatted }</td>
        </tr>
      )
    });
  }

  const clientsOptions = () => {
    return clientNames.map((clientName) => 
      <option key={clientName} value={clientName}>{clientName}</option>
    );
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 bg-slate-950 w-600">
      <div className="align-items-center">
        <Image 
          src="/Logo_C6_Bank.png" 
          alt="Vercel Logo"
          className="dark:invert"
          width={140}
          height={48}
          priority
        />
        <div className="mt-5">
          <form method="post" id='bill-upload-form'>
            <input type="file" name="bill_file" id="bill-file-selector" onChange={($event) => UploadFile($event.target)}></input>
          </form>
        </div>

        <div className="mt-6 text-left">
          { 
            purchasesData.length > 0 ? 
              (<div className="mb-6">
                <select className="bg-gray-800 py-2 px-1 rounded" onChange={filterPurchases}>
                  <option value="" selected>Todos os clientes</option>
                  { clientsOptions() }
                </select>
              </div>
              )
            : <></>
          }

          <table className="table-auto">
            <thead className='bg-slate-50 dark:bg-slate-700'>
              <tr>
                <th className='border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left'>Data</th>
                <th className='border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left'>Cliente</th>
                <th className='border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left'>Descrição</th>
                <th className='border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left'>Categoria</th>
                <th className='border border-slate-300 dark:border-slate-600 font-semibold p-4 text-slate-900 dark:text-slate-200 text-left'>Valor R$</th>
              </tr>
            </thead>
            <tbody>
              { renderItems() }
            </tbody>
          </table>
        </div> 
      </div>
    </main>
  )
}
