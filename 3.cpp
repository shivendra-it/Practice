#include<bits/stdc++.h>
using namespace std;

int cmp(const void *a,const void *b) {
    return ((const int *)a)[0] - ((const int *)b)[0];
}

int main(){
int t;
cin>>t;
while(t--){
int ply;
cin>>ply;
int arr1[105];
fill(arr1,arr1+105,0);
map<int,map< int, int > > mp;
for(int i=0;i<ply;i++){
  int arr[10][2]={0};
  int cook;
  cin>>cook;
  for(int j=0;j<cook;j++){
    int c;
    cin>>c;
    mp[i][c-1]++;
}

for (int l=0;l<6;l++){
  arr[l][0]=mp[i][l];
  arr[l][1]=l;
}

qsort(arr,6,2*sizeof(int),cmp);

int cost =cook;


for(int l=0;l<6;l++){
  if(arr[l][0]==0){
    continue;
  }else{
    if(l==0){
      cost += 4*arr[l][0];
      arr[]
    }
    if(l==1)


  }

}


arr1[i]=cost;
}
int flag =-1,ind = -1,max=-1;

for(int k=0;k<ply;k++){
  if(arr1[k]>max){
    max =arr1[k];
    ind = k;
  }else if(arr1[k]==max){
    flag=2;
  }
}

if(flag==2){
  cout<<"tie"<<endl;
}else{
  if(ind==0){
    cout<<"chef"<<endl;
  }else
  {
    cout<<ind+1<<endl;
  }
}

}
return 0;
}
