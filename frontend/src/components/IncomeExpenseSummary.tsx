'use client';

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
  const incomeItems = [
    { title: '今月', amount: '100,000円' },
    { title: '今週', amount: '100,000円' },
    { title: '今日', amount: '100,000円' },
  ];

  const expenseItems = [
    { title: '今月', amount: '100,000円' },
    { title: '今週', amount: '100,000円' },
    { title: '今日', amount: '100,000円' },
  ];

  return (
    <div className="rounded-lg bg-white p-6 shadow-md">
      <h1 className="mb-4 text-2xl font-bold">収支集計</h1>
      <SummaryRow category="収入" items={incomeItems} />
      <SummaryRow category="支出" items={expenseItems} />
      <div className="flex items-center justify-end border-t pt-4">
        <h2 className="mr-4 text-lg font-bold">残高</h2>
        <p className="font-semibold">100,000円</p>
      </div>
    </div>
  );
}
