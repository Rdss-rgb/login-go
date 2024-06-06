var tabs=document.querySelectorAll('[data-pane="tabs"]')
var tabs_child=document.querySelectorAll('[data-pane="tabs-child"]')
var p=document.getElementById('name')



function switchpane(num){
            for(var i=0; i < tabs.length; i++){
    
         
                if(i==num){
           
                    tabs[i].classList.add('border-orange-500','text-black')
                    tabs[i].classList.remove('hover:text-gray-600' ,'hover:border-gray-300','border-transparent')

                    for (let j = 0; j < tabs_child.length; j++) {
                        var shifttotheleft=tabs_child[j].id
                        if(shifttotheleft==num+1){
                            tabs_child[j].classList.remove('hidden')
                            console.log('shifttothebyte')
                        }
                        else{
                            console.log('shifttothebits')
                            tabs_child[j].classList.add('hidden')
                        }
                        
                    }
                }
                else{
                    tabs[i].classList.remove('border-orange-500','text-black')
                    tabs[i].classList.add('hover:text-gray-600' ,'hover:border-gray-300','border-transparent')
                }
                    
            }

}
switchpane(0)




// var main_body=document.getElementById('main-body')



// var base_body=document.createElement('div')
// base_body.setAttribute('data','data-pane="tabs-child"')
// base_body.innerHTML=`    {{range .FoodList}}
//     <div class="w-64 h-64 bg-slate-100 mt-5 rounded-md overflow-hidden relative hover:translate-y-2 transition-transform">
//         <img src="{{.Icon}}" alt="" srcset="" class="rounded-lg max-h-36 min-h-36 w-full">
//         <div class="p-4">
//             <p id="name" class="text-lg font-bold">{{.Name}}</p>
//             <div class="flex gap-2">
//                 <span class="font-bold text-lg font-medium"><span class="text-orange-500">₱</span>{{.Price}}</span>
//                 <span class="line-through text-slate-600 font-light text-lg">₱ {{.Discount}}</span>
//             </div>
//             <button class="absolute right-2 bottom-2 bg-orange-500 text-white px-4 font-light text-lg rounded-full hover:bg-orange-600"> order{{.Category}}</button>
//         </div>
//     </div>
//     {{end}}
//  `;

// main_body.appendChild(base_body)