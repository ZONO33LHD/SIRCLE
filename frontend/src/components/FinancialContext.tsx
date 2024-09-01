'use client';

import { gql, useQuery } from '@apollo/client';
import React, { createContext, useContext, ReactNode } from 'react';
import { format } from 'date-fns';

export const GET_INCOME_EXPENSE_SUMMARY = gql`
  query GetIncomeExpenseSummary($startDate: String!, $endDate: String!) {
    getIncomeExpenseSummary(startDate: $startDate, endDate: $endDate) {
      incomeItems {
        title
        amount
      }
      expenseItems {
        title
        amount
      }
      balance
    }
  }
`;

type SummaryItemProps = {
  title: string;
  amount: string;
};

type FinancialData = {
  incomeItems: SummaryItemProps[];
  expenseItems: SummaryItemProps[];
  balance: string;
};

type DateRange = {
  startDate: Date;
  endDate: Date;
};

interface FinancialContextType {
  getFinancialDataForPeriod: (range: DateRange) => Promise<FinancialData>;
  updateFinancialData: () => void;
}

const FinancialContext = createContext<FinancialContextType | undefined>(undefined);

type FinancialProviderProps = {
  children: ReactNode;
};

export const FinancialProvider: React.FC<FinancialProviderProps> = ({ children }) => {
  const { refetch: getIncomeExpenseSummaryQuery } = useQuery(GET_INCOME_EXPENSE_SUMMARY, {
    skip: true,
    variables: {
      startDate: new Date().toISOString(),
      endDate: new Date().toISOString(),
    },
  });

  const getFinancialDataForPeriod = async (range: DateRange): Promise<FinancialData> => {
    try {
      const formattedStartDate = format(range.startDate, "yyyy-MM-dd'T'00:00:00+09:00");
      const formattedEndDate = format(range.endDate, "yyyy-MM-dd'T'23:59:59+09:00");
      console.log('Sending query with variables:', {
        startDate: formattedStartDate,
        endDate: formattedEndDate,
      });
      const { data } = await getIncomeExpenseSummaryQuery({
        variables: {
          startDate: formattedStartDate,
          endDate: formattedEndDate,
        },
      });
      if (data && data.getIncomeExpenseSummary) {
        return {
          incomeItems: data.getIncomeExpenseSummary.incomeItems.map((item: any) => ({
            title: item.title,
            amount: `${item.amount}円`,
          })),
          expenseItems: data.getIncomeExpenseSummary.expenseItems.map((item: any) => ({
            title: item.title,
            amount: `${item.amount}円`,
          })),
          balance: `${data.getIncomeExpenseSummary.balance}円`,
        };
      }
      throw new Error('データが取得できませんでした');
    } catch (error) {
      console.error('データの取得に失敗しました:', error);
      throw error;
    }
  };

  return (
    <FinancialContext.Provider
      value={{
        getFinancialDataForPeriod,
        updateFinancialData: () => {},
      }}
    >
      {children}
    </FinancialContext.Provider>
  );
};

export const useFinancial = () => {
  const context = useContext(FinancialContext);
  if (context === undefined) {
    throw new Error('useFinancial must be used within a FinancialProvider');
  }
  return context;
};
