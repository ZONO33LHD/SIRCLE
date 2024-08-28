'use client';

import { gql, useMutation } from '@apollo/client';
import React, { createContext, useContext, ReactNode } from 'react';

const GET_INCOME_EXPENSE_SUMMARY = gql`
mutation GetIncomeExpenseSummary($startDate: String!, $endDate: String!) {
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
  const [getIncomeExpenseSummary] = useMutation(GET_INCOME_EXPENSE_SUMMARY);

  const getFinancialDataForPeriod = async (range: DateRange): Promise<FinancialData> => {
    try {
      const { data } = await getIncomeExpenseSummary({
        variables: {
          period: 'custom',
          date: range.endDate.toISOString(),
        },
      });
      if (data && data.getIncomeExpenseSummary) {
        return data.getIncomeExpenseSummary;
      }
    } catch (error) {
      console.error('データの取得に失敗しました:', error);
    }
    return { incomeItems: [], expenseItems: [], balance: '0円' };
  };

  return (
    <FinancialContext.Provider value={{
      getFinancialDataForPeriod,
      updateFinancialData: () => {},
    }}>
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