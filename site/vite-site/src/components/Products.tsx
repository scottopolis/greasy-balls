import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { getApiUrl } from "@/util/getApiUrl";
import { useQuery } from "@tanstack/react-query";

interface Product {
  id: number;
  name: string;
  price: number;
  image?: string;
}

const demoProducts: Product[] = [
  {
    id: 1,
    name: "Leather Wallet",
    price: 59.99,
    image: "picsum.photos/200/200",
  },
  {
    id: 2,
    name: "Wireless Earbuds",
    price: 129.99,
    image: "picsum.photos/200/200",
  },
  {
    id: 3,
    name: "Smartwatch",
    price: 199.99,
    image: "picsum.photos/200/200",
  },
  {
    id: 4,
    name: "Portable Charger",
    price: 39.99,
    image: "picsum.photos/200/200",
  },
  {
    id: 5,
    name: "Sunglasses",
    price: 79.99,
    image: "picsum.photos/200/200",
  },
  {
    id: 6,
    name: "Backpack",
    price: 89.99,
    image: "picsum.photos/200/200",
  },
];

export function Products() {
  async function getProducts() {
    const apiUrl = getApiUrl();
    const response = fetch(`${apiUrl}/products`).then((res) => res.json());

    return response;
  }

  const getProductsQuery = useQuery({
    queryKey: ["products"],
    queryFn: async () => await getProducts(),
  });

  const products: any = getProductsQuery?.data || demoProducts;

  console.log(products);

  return (
    <>
      <h1 className="text-3xl font-bold mb-6">Our Products</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {!getProductsQuery?.isFetching &&
          products?.map((product: any) => (
            <Card key={product.id} className="overflow-hidden">
              <CardContent className="p-4">
                {product.image ? (
                  <img
                    src={`https://${product.image}`}
                    alt={product.name}
                    width={200}
                    height={200}
                    className="w-full h-48 object-cover mb-4 rounded-md"
                  />
                ) : null}
                <h2 className="text-lg font-semibold mb-2">{product.name}</h2>
                <p className="text-gray-600">${product.price.toFixed(2)}</p>
              </CardContent>
              <CardFooter>
                <Button className="w-full">Add to Cart</Button>
              </CardFooter>
            </Card>
          ))}
      </div>
    </>
  );
}
