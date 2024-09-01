'use client';

import React, { useState, useEffect } from 'react';
import { Calendar } from '@/components/ui/calendar';
import { format, differenceInDays, differenceInMonths, differenceInYears } from 'date-fns';
import { ja } from 'date-fns/locale';
import { Button } from '@/components/ui/button';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { useFinancial } from './FinancialContext';

type SummaryItemProps = {
  title: string;
  amount: string;
};

const SummaryItem = ({ title, amount }: SummaryItemProps) => (
  <div>
    <h4 className="text-sm text-gray-600">{title}</h4>
    <p className="font-semibold">{amount}</p>
  </div>
);

const SummaryRow = ({ category, items }: { category: string; items: SummaryItemProps[] }) => (
  <div className="flex items-center justify-between py-4">
    <h2 className="w-1/4 text-lg font-bold">{category}</h2>
    {items.map((item, index) => (
      <SummaryItem key={index} {...item} />
    ))}
  </div>
);

export default function IncomeExpenseSummary() {
  const currentYear = new Date().getFullYear();
  const [startDate, setStartDate] = useState<Date>(new Date(currentYear, 0, 1));
  const [endDate, setEndDate] = useState<Date>(new Date(currentYear, 11, 31));
  const [isStartDateOpen, setIsStartDateOpen] = useState(false);
  const [isEndDateOpen, setIsEndDateOpen] = useState(false);
  const [periodData, setPeriodData] = useState<{
    incomeItems: SummaryItemProps[];
    expenseItems: SummaryItemProps[];
    balance: string;
  }>({ incomeItems: [], expenseItems: [], balance: '0円' });
  const { getFinancialDataForPeriod } = useFinancial();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getFinancialDataForPeriod({
          startDate: startDate,
          endDate: endDate
        });
        setPeriodData(data);
      } catch (error) {
        console.error('データの取得に失敗しました:', error);
      }
    };
    fetchData();
  }, [startDate, endDate, getFinancialDataForPeriod]);

  const formatPeriod = (start: Date, end: Date) => {
    const days = differenceInDays(end, start);
    const months = differenceInMonths(end, start);
    const years = differenceInYears(end, start);

    if (years > 0) {
      return `${years}年と${months % 12}ヶ月`;
    } else if (months > 0) {
      return `${months}ヶ月と${days % 30}日`;
    } else {
      return `${days + 1}日間`;
    }
  };

  const renderDateSelector = (
    date: Date,
    setDate: (date: Date) => void,
    isOpen: boolean,
    setIsOpen: (isOpen: boolean) => void
  ) => (
    <Popover open={isOpen} onOpenChange={setIsOpen}>
      <PopoverTrigger asChild>
        <Button variant="outline">{format(date, 'yyyy/MM/dd', { locale: ja })}</Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0">
        <Calendar
          mode="single"
          selected={date}
          onSelect={(date: Date | undefined) => {
            if (date) {
              setDate(date);
              setIsOpen(false);
            }
          }}
          locale={ja}
        />
      </PopoverContent>
    </Popover>
  );

  return (
    <div className="rounded-lg bg-white p-6 shadow-md">
      <div className="mb-6 flex items-center justify-between">
        <h1 className="text-2xl font-bold">収支集計</h1>
        <div className="flex items-center space-x-4">
          {renderDateSelector(startDate, setStartDate, isStartDateOpen, setIsStartDateOpen)}
          <span>〜</span>
          {renderDateSelector(endDate, setEndDate, isEndDateOpen, setIsEndDateOpen)}
        </div>
      </div>
      <div className="mb-4">
        <p>期 間： {formatPeriod(startDate, endDate)}</p>
      </div>
      <SummaryRow category="収入" items={periodData.incomeItems} />
      <SummaryRow category="支出" items={periodData.expenseItems} />
      <div className="flex items-center justify-end border-t pt-4">
        <h2 className="mr-4 text-lg font-bold">残高</h2>
        <p className="font-semibold">{periodData.balance}</p>
      </div>
    </div>
  );
}
